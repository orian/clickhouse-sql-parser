
        CREATE TABLE promql_row_cost_ts ENGINE = TimeSeries
        SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp)
        TAGS INNER COLUMNS (id Tuple(UInt64, UUID))
    