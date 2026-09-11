import json
import subprocess
import tempfile
import unittest
from pathlib import Path
from extract_clickhouse_sql import extract, xml_queries


class ClickHouseExtractionTests(unittest.TestCase):
    def test_xml_text_cdata_and_metadata(self):
        examples = xml_queries(b'''<test>
<query>SELECT 1 &lt; 2</query>
<create_query><![CDATA[CREATE TABLE t (x String DEFAULT ';') ENGINE=Memory]]></create_query>
<query runs="1">SELECT {n}</query>
<substitutions><substitution><name>n</name><values><value>1</value></values></substitution></substitutions>
</test>''')
        self.assertEqual([x[0] for x in examples], [b'SELECT 1 < 2', b"CREATE TABLE t (x String DEFAULT ';') ENGINE=Memory", b'SELECT {n}'])
        self.assertEqual(examples[0][1]['line'], 2)
        self.assertEqual(examples[2][1]['attributes'], {'runs': '1'})
        self.assertTrue(all(x[1]['has_substitutions'] for x in examples))

    def test_copy_inventory_and_repeatability(self):
        with tempfile.TemporaryDirectory() as tmp:
            source, output = Path(tmp) / 'source', Path(tmp) / 'corpus'
            source.mkdir()
            subprocess.run(['git', 'init', '-q', str(source)], check=True)
            subprocess.run(['git', '-C', str(source), '-c', 'user.name=Test', '-c', 'user.email=test@example.com', '-c', 'commit.gpgsign=false', 'commit', '--allow-empty', '-qm', 'fixture'], check=True)
            original = b"SELECT '\xff';\r\nSELECT (; -- { serverError SYNTAX_ERROR }\r\n"
            (source / 'one.sql').write_bytes(original)
            (source / 'two.sql').write_bytes(original)
            (source / 'template.sql.j2').write_text('{% for n in range(2) %}SELECT {{n}};{% endfor %}')
            extract(source, output)
            before = (output / 'manifest.json').read_bytes()
            manifest = json.loads(before)
            self.assertEqual(manifest['occurrences'], 2)
            self.assertEqual(len(manifest['examples']), 1)
            self.assertEqual(manifest['unrendered_templates'], ['template.sql.j2'])
            example = manifest['examples'][0]
            self.assertEqual((output / example['file']).read_bytes(), original)
            self.assertEqual(example['sources'][0]['expected_errors'], ['{ serverError SYNTAX_ERROR }'])
            extract(source, output)
            self.assertEqual(before, (output / 'manifest.json').read_bytes())


if __name__ == '__main__':
    unittest.main()
