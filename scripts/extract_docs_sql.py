#!/usr/bin/env python3
"""Extract explicitly SQL-marked documentation examples without rewriting SQL."""
import argparse
import hashlib
import json
import re
from pathlib import Path

# SQL and ClickHouse plus dialect labels used by integrations in the docs.
LANGUAGES = {'sql', 'clickhouse', 'clickhouse-sql', 'chsql', 'pgsql', 'postgresql', 'mysql', 'tsql'}
FENCE = re.compile(r'^\s*(`{3,}|~{3,})(.*)$')


def blocks(text):
    """Respect non-SQL outer fences; support MDX indentation and fence metadata."""
    opened = None
    body = []
    for number, line in enumerate(text.splitlines(keepends=True), 1):
        match = FENCE.match(line.rstrip('\r\n'))
        if opened is None:
            if match:
                marker, info = match.groups()
                opened = (marker, info.strip(), number)
                body = []
        elif match and match[1][0] == opened[0][0] and len(match[1]) >= len(opened[0]) and not match[2].strip():
            marker, info, start = opened
            language = info.split()[0].lower() if info else ''
            if language in LANGUAGES:
                yield ''.join(body), {'line': start + 1, 'end_line': number - 1,
                                      'info': info, 'kind': 'fence'}
            opened = None
        else:
            body.append(line)
    if opened:
        info = opened[1]
        if info and info.split()[0].lower() in LANGUAGES:
            raise ValueError(f'unclosed SQL fence at line {opened[2]}')



def component_blocks(text):
    """Resolve the static template maps used by SQL CodeBlocks in this snapshot.

    Never execute documentation JavaScript; fail on unhandled SQL components so
    future documentation changes cannot silently reduce extraction coverage.
    """
    for component in re.finditer(r'<CodeBlock\s+language=["\']sql["\'][^>]*>(.*?)</CodeBlock>', text, re.S | re.I):
        expression = component[1].strip()
        variable = re.fullmatch(r'\{(\w+)\}', expression)
        templates = []
        if variable:
            binding = re.search(r'const\s+' + re.escape(variable[1]) + r'\s*=\s*(\w+)\[', text)
            if binding:
                mapping = re.search(r'const\s+' + re.escape(binding[1]) + r'\s*=\s*\{(.*?)\n\s*\};', text, re.S)
                if mapping:
                    templates = [(m[1], mapping.start(1) + m.start(1)) for m in re.finditer(r'`([^`]*?)`', mapping[1], re.S)]
        else:
            literal = re.fullmatch(r'\{`([^`]*)`\}', expression, re.S)
            if literal:
                templates = [(literal[1], component.start(1))]
        if not templates or any('${' in sql or '\\' in sql for sql, _ in templates):
            raise ValueError(f'unresolved SQL CodeBlock at line {text[:component.start()].count(chr(10)) + 1}')
        for sql, start in templates:
            line = text[:start].count('\n') + 1
            yield sql, {'line': line, 'end_line': line + sql.count('\n'),
                        'info': 'sql (static MDX template)', 'kind': 'component'}

def extract(source, output):
    entries = {}
    skipped = []
    files = 0
    for path in sorted(source.rglob('*')):
        if any(part.startswith('.') for part in path.relative_to(source).parts):
            continue
        if path.suffix.lower() not in {'.md', '.mdx', '.sql'}:
            continue
        if not path.is_file():
            if path.is_symlink():
                skipped.append(path.relative_to(source).as_posix())
            continue
        files += 1
        text = path.read_bytes().decode('utf-8')
        try:
            examples = [(text, {'line': 1, 'end_line': len(text.splitlines()), 'info': 'sql', 'kind': 'file'})] if path.suffix.lower() == '.sql' else [*blocks(text), *component_blocks(text)]
        except ValueError as error:
            raise ValueError(f'{path}: {error}') from error
        for sql, location in examples:
            # Only byte-identical blocks are deduplicated, including translations.
            digest = hashlib.sha256(sql.encode('utf-8')).hexdigest()
            entry = entries.setdefault(digest, {'id': digest, 'file': f'queries/{digest}.sql', 'sources': []})
            entry['sources'].append({'path': path.relative_to(source).as_posix(), **location})
            target = output / entry['file']
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_bytes(sql.encode('utf-8'))
    if not entries:
        raise ValueError(f'no SQL examples found in {source}')
    # Delete only obsolete content-addressed files owned by this generator.
    for path in (output / 'queries').glob('*.sql'):
        if re.fullmatch('[0-9a-f]{64}', path.stem) and path.stem not in entries:
            path.unlink()
    manifest = {'version': 1, 'languages': sorted(LANGUAGES), 'scanned_files': files,
                'skipped_broken_symlinks': skipped,
                'occurrences': sum(len(e['sources']) for e in entries.values()),
                'examples': [entries[key] for key in sorted(entries)]}
    (output / 'manifest.json').write_text(json.dumps(manifest, ensure_ascii=False, indent=2) + '\n')
    print(f"{files} source files; {manifest['occurrences']} SQL occurrences; {len(entries)} unique examples")
    if skipped:
        print(f'{len(skipped)} broken links recorded in manifest')


if __name__ == '__main__':
    root = Path(__file__).resolve().parents[1]
    cli = argparse.ArgumentParser(description=__doc__)
    cli.add_argument('--source', type=Path, default=root / 'clickhouse-docs')
    cli.add_argument('--output', type=Path, default=root / 'docs-sql')
    args = cli.parse_args()
    extract(args.source, args.output)
