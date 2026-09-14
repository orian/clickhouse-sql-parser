CREATE TABLE my_table (...)
ENGINE = MergeTree()
ORDER BY (timestamp, id)
SETTINGS 
    -- Увеличить максимальное количество потоков вставки для параллельной записи частей
    max_insert_threads = 4,
    -- Разрешить вставки с кворумом для надёжности (ReplicatedMergeTree)
    insert_quorum = 2
