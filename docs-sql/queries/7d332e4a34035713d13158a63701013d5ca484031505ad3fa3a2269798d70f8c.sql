    CREATE TABLE my_clickhouse_table
    (
        -- defina as colunas para corresponder à sua tabela Iceberg
        `id` Int64,
        `name` String,
        `event_time` DateTime64(3)
    )
    ENGINE = MergeTree
    ORDER BY id;

    INSERT INTO my_clickhouse_table
    SELECT * FROM polaris_catalog.`polaris_db.my_iceberg_table`;
