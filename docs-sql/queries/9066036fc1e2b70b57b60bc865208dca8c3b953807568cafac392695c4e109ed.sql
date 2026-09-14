CREATE TABLE my_table (...)
ENGINE = MergeTree()
ORDER BY (timestamp, id)
SETTINGS 
    -- Aumentar el máximo de hilos de inserción para escritura paralela de partes
    max_insert_threads = 4,
    -- Permitir inserciones con quórum para mayor confiabilidad (ReplicatedMergeTree)
    insert_quorum = 2
