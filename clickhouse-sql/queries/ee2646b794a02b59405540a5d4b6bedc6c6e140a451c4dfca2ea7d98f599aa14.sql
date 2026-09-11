
        CREATE TABLE fact_1000p (p UInt16, id UInt64)
        ENGINE = MergeTree PARTITION BY p ORDER BY id
        SETTINGS refresh_statistics_interval = 0
    