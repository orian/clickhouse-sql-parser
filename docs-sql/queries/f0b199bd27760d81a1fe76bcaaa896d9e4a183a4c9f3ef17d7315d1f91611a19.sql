SET allow_experimental_qbit_type = 1;

-- Assumindo que você tem uma tabela com embeddings Float32
ALTER TABLE dbpedia ADD COLUMN qbit QBit(Float32, 1536);
ALTER TABLE dbpedia UPDATE qbit = vector WHERE 1;
