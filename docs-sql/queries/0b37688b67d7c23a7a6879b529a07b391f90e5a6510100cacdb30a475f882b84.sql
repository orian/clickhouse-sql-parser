    CREATE TABLE my_clickhouse_table
    (
        -- Iceberg 테이블에 맞게 컬럼을 정의합니다
        `id` Int64,
        `name` String,
        `event_time` DateTime64(3)
    )
    ENGINE = MergeTree
    ORDER BY id;

    INSERT INTO my_clickhouse_table
    SELECT * FROM polaris_catalog.`polaris_db.my_iceberg_table`;
