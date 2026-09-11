import tempfile
import unittest
from pathlib import Path
from extract_docs_sql import blocks, component_blocks, extract
import json


class ExtractionTests(unittest.TestCase):
    def test_fences(self):
        text = '````markdown\n```sql\nnot SQL\n```\n````\n  ~~~SQL title="Query"\n  SELECT 1;\n  ~~~\n```response\n42\n```\n```sql\nSELECT 2;\n```\n'
        result = list(blocks(text))
        self.assertEqual([sql for sql, _ in result], ['  SELECT 1;\n', 'SELECT 2;\n'])
        self.assertEqual(result[0][1]['line'], 7)

    def test_unclosed_fence(self):
        with self.assertRaises(ValueError):
            list(blocks('```sql\nSELECT 1;'))

    def test_component_map(self):
        text = 'const variants = {\n a: `SELECT 1;`,\n b: `SELECT 2;`,\n};\nconst query = variants[engine];\n<CodeBlock language="sql">{query}</CodeBlock>'
        self.assertEqual([s for s, _ in component_blocks(text)], ['SELECT 1;', 'SELECT 2;'])
        with self.assertRaises(ValueError):
            list(component_blocks('<CodeBlock language="sql">{unknown}</CodeBlock>'))

    def test_deduplication_provenance_and_repeatability(self):
        with tempfile.TemporaryDirectory() as temp:
            root = Path(temp)
            source, output = root / 'docs', root / 'out'
            source.mkdir()
            (source / 'a.mdx').write_text('```sql\nSELECT 1;\n```\n```sql\nSELECT 1;\n```\n')
            (source / 'b.sql').write_text('SELECT 1;\n')
            (source / 'broken.md').symlink_to('missing')
            extract(source, output)
            original = (output / 'manifest.json').read_bytes()
            manifest = json.loads(original)
            self.assertEqual(manifest['occurrences'], 3)
            self.assertEqual(len(manifest['examples']), 1)
            self.assertEqual(manifest['skipped_broken_symlinks'], ['broken.md'])
            extract(source, output)
            self.assertEqual(original, (output / 'manifest.json').read_bytes())


if __name__ == '__main__':
    unittest.main()
