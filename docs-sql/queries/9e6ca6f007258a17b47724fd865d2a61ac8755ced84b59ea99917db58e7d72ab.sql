SET allow_experimental_qbit_type = 1;

-- Suponiendo que tienes una tabla con embeddings Float32
ALTER TABLE dbpedia ADD COLUMN qbit QBit(Float32, 1536);
ALTER TABLE dbpedia UPDATE qbit = vector WHERE 1;
