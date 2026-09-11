SET allow_experimental_qbit_type = 1;

-- Float32 임베딩이 포함된 테이블이 있다고 가정합니다
ALTER TABLE dbpedia ADD COLUMN qbit QBit(Float32, 1536);
ALTER TABLE dbpedia UPDATE qbit = vector WHERE 1;
