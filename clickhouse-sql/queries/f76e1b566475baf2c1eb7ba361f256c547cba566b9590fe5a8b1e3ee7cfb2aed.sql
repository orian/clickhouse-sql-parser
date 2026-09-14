
        CREATE TABLE promql_pk_range_ts ENGINE = TimeSeries
        SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity = 64
        TAGS INNER COLUMNS (id Tuple(UInt64, UUID))
    