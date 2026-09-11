
        SETTINGS="--max_threads=8 --max_insert_threads=8 --max_threads_min_free_memory_per_thread=0 --max_insert_threads_min_free_memory_per_thread=0"
        $CLICKHOUSE_LOCAL $SETTINGS --query "
            SELECT number AS id, rand(1) % 1000000 AS val1, rand(2) % 1000000 AS val2,
                   concat('item_', toString(rand(3) % 100000)) AS name, rand(4) / 1000.0 AS amount,
                   toDateTime('2020-01-01 00:00:00') + (rand(5) % 100000000) AS ts
            FROM numbers(10000000) FORMAT Native" \
        | $CLICKHOUSE_LOCAL $SETTINGS --query "
            CREATE TABLE t
            (
                id UInt64 CODEC(ZSTD(3)),
                val1 UInt32 CODEC(ZSTD(3)),
                val2 UInt32 CODEC(ZSTD(3)),
                name String CODEC(ZSTD(3)),
                amount Float64 CODEC(ZSTD(3)),
                ts DateTime CODEC(ZSTD(3))
            )
            ENGINE = MergeTree ORDER BY (val1, val2, name);
            INSERT INTO t FORMAT Native"
    