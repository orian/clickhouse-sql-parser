
        CREATE TABLE dim_j (id UInt64)
        ENGINE = MergeTree ORDER BY id
        SETTINGS refresh_statistics_interval = 0
    