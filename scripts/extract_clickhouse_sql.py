#!/usr/bin/env python3
"""Copy SQL files and extract performance XML queries from a ClickHouse checkout."""
import argparse
import hashlib
import json
import re
import subprocess
from pathlib import Path
from xml.parsers import expat

QUERY_TAGS = {'query', 'create_query', 'fill_query', 'drop_query'}


def xml_queries(data):
    """Decode XML text/CDATA only; never substitute or rewrite SQL."""
    parser = expat.ParserCreate()
    queries = []
    active = None
    substitutions = False

    def start(tag, attrs):
        nonlocal active, substitutions
        if tag == 'substitutions':
            substitutions = True
        if tag in QUERY_TAGS:
            if active is not None:
                raise ValueError('nested SQL query elements')
            active = [tag, parser.CurrentLineNumber, [], attrs]

    def characters(text):
        if active is not None:
            active[2].append(text)

    def end(tag):
        nonlocal active
        if active is not None and tag == active[0]:
            kind, line, pieces, attrs = active
            queries.append((''.join(pieces).encode('utf-8'), {
                'line': line, 'end_line': parser.CurrentLineNumber,
                'info': tag, 'kind': 'performance_xml', 'attributes': attrs,
            }))
            active = None

    parser.StartElementHandler = start
    parser.CharacterDataHandler = characters
    parser.EndElementHandler = end
    parser.Parse(data, True)
    for _, location in queries:
        location['has_substitutions'] = substitutions
    return queries


def extract(source, output):
    if not source.is_dir():
        raise ValueError(f'missing ClickHouse checkout: {source}')
    entries, templates, broken = {}, [], []
    counts = {'sql_files': 0, 'performance_xml_files': 0}
    revision = subprocess.run(['git', '-C', str(source), 'rev-parse', 'HEAD'],
                              check=True, capture_output=True, text=True).stdout.strip()
    for path in sorted(source.rglob('*')):
        relative = path.relative_to(source)
        if any(part.startswith('.') for part in relative.parts):
            continue
        if path.name.endswith(('.sql.j2', '.sql.in')):
            templates.append(relative.as_posix())
            continue
        is_xml = path.suffix == '.xml' and relative.parts[:2] in {
            ('tests', 'performance'), ('tests', 'perf_drafts')}
        if path.suffix != '.sql' and not is_xml:
            continue
        if not path.is_file():
            if path.is_symlink():
                broken.append(relative.as_posix())
            continue
        data = path.read_bytes()
        if is_xml:
            try:
                examples = xml_queries(data)
            except (expat.ExpatError, ValueError) as error:
                raise ValueError(f'{relative}: {error}') from error
            counts['performance_xml_files'] += 1
        else:
            counts['sql_files'] += 1
            examples = [(data, {'line': 1, 'end_line': len(data.splitlines()),
                                'info': 'sql', 'kind': 'file'})]
        for sql, location in examples:
            # Annotations are retained verbatim in SQL and exposed as metadata.
            annotations = re.findall(rb'\{\s*(?:serverError|clientError)\s+[^}\r\n]*\}', sql)
            if annotations:
                location['expected_errors'] = sorted({a.decode('utf-8', errors='replace') for a in annotations})
            digest = hashlib.sha256(sql).hexdigest()
            entry = entries.setdefault(digest, {'id': digest, 'file': f'queries/{digest}.sql', 'sources': []})
            entry['sources'].append({'path': relative.as_posix(), **location})
            target = output / entry['file']
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_bytes(sql)
    if not entries:
        raise ValueError('no SQL examples found')
    for path in (output / 'queries').glob('*.sql'):
        if re.fullmatch('[0-9a-f]{64}', path.stem) and path.stem not in entries:
            path.unlink()
    manifest = {'version': 1, 'source_revision': revision, 'scanned_files': counts,
                'unrendered_templates': templates, 'skipped_broken_symlinks': broken,
                'occurrences': sum(len(e['sources']) for e in entries.values()),
                'examples': [entries[key] for key in sorted(entries)]}
    (output / 'manifest.json').write_text(json.dumps(manifest, ensure_ascii=False, indent=2) + '\n')
    print(f"{counts}; {manifest['occurrences']} occurrences; {len(entries)} unique examples; {len(templates)} unrendered templates")


if __name__ == '__main__':
    root = Path(__file__).resolve().parents[1]
    cli = argparse.ArgumentParser(description=__doc__)
    cli.add_argument('--source', type=Path, default=root.parent / 'ClickHouse')
    cli.add_argument('--output', type=Path, default=root / 'clickhouse-sql')
    args = cli.parse_args()
    extract(args.source, args.output)
