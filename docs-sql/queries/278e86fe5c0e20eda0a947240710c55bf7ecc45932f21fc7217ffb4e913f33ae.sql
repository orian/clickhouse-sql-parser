CREATE TABLE my_table (...)
ENGINE = MergeTree()
ORDER BY (timestamp, id)
SETTINGS 
    -- 병렬 파트 쓰기를 위한 최대 삽입 스레드 수 증가
    max_insert_threads = 4,
    -- 안정성을 위해 쿼럼(quorum) 삽입 허용 (ReplicatedMergeTree)
    insert_quorum = 2
