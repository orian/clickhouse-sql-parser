
        CREATE TABLE ts_bucketing (id UInt64, timestamp DateTime64(3), value Float64) ENGINE = MergeTree ORDER BY (id, timestamp)
    