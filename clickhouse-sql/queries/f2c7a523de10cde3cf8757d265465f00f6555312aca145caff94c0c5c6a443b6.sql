
        CREATE TABLE ts_scalar_add (id UInt16, timestamp DateTime64(3), value Float64) ENGINE = MergeTree ORDER BY (id, timestamp)
    