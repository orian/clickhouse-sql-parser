
        CREATE TABLE ts_grid_sparse (id UInt16, timestamp DateTime, value Float64) ENGINE = MergeTree ORDER BY (id, timestamp)
    