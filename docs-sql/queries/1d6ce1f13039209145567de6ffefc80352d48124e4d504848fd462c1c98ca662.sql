SET allow_experimental_qbit_type = 1;

-- 假设您有一个包含 Float32 嵌入向量的表
ALTER TABLE dbpedia ADD COLUMN qbit QBit(Float32, 1536);
ALTER TABLE dbpedia UPDATE qbit = vector WHERE 1;
