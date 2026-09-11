SET allow_experimental_qbit_type = 1;

-- Assuming you have a table with Float32 embeddings
ALTER TABLE dbpedia ADD COLUMN qbit QBit(Float32, 1536);
ALTER TABLE dbpedia UPDATE qbit = vector WHERE 1;
