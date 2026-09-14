CREATE TABLE my_table (...)
ENGINE = MergeTree()
ORDER BY (timestamp, id)
SETTINGS 
    -- Aumentar o número máximo de insert threads para escrita paralela de partes
    max_insert_threads = 4,
    -- Permitir inserts com quorum para maior confiabilidade (ReplicatedMergeTree)
    insert_quorum = 2
