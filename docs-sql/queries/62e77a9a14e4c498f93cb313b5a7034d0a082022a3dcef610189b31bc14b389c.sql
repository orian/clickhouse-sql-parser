CREATE TABLE my_table (...)
ENGINE = MergeTree()
ORDER BY (timestamp, id)
SETTINGS 
    -- 增加最大插入线程数以实现并行分片写入
    max_insert_threads = 4,
    -- 允许使用 quorum 写入以提高可靠性（ReplicatedMergeTree）
    insert_quorum = 2
