
        CREATE MATERIALIZED VIEW hits_mv ENGINE MergeTree
        PARTITION BY toYYYYMM(EventDate)
        ORDER BY (CounterID, EventDate, intHash32(UserID))
        SAMPLE BY intHash32(UserID)
        SETTINGS
            parts_to_delay_insert = 5000,
            parts_to_throw_insert = 5000
        AS
            -- don't select all columns to keep the run time down
            SELECT CounterID, EventDate, UserID, Title
            FROM hits_10m_single
            -- do not select anything because we only need column types
            LIMIT 0
    