SET allow_experimental_qbit_type = 1;

-- Float32の埋め込みを持つテーブルがある場合
ALTER TABLE dbpedia ADD COLUMN qbit QBit(Float32, 1536);
ALTER TABLE dbpedia UPDATE qbit = vector WHERE 1;
