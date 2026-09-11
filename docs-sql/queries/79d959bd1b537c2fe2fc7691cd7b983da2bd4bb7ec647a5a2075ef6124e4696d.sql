    CREATE DATABASE IF NOT EXISTS test ON CLUSTER cluster_2S_1R;
    CREATE TABLE test.test_table ON CLUSTER cluster_2S_1R
    (
        `id` UInt64,
        `name` String
    )
    ENGINE = MergeTree()
    ORDER BY id;
