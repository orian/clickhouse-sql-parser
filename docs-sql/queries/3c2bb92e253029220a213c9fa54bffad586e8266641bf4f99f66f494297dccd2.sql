CREATE TABLE my_table (...)
ENGINE = MergeTree()
ORDER BY (timestamp, id)
SETTINGS 
    -- 並列パート書き込みのために挿入スレッドの最大数を増やす
    max_insert_threads = 4,
    -- 信頼性向上のためにクォーラムを使用した挿入を許可する（ReplicatedMergeTree）
    insert_quorum = 2
