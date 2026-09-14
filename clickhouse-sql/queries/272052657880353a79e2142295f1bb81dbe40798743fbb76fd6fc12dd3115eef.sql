
        CREATE TABLE ts_bucketing_unsorted (id UInt64, h UInt64, timestamp DateTime64(3), value Float64) ENGINE = MergeTree ORDER BY (id, h)
    