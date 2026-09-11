
        CREATE TABLE hits_wide AS hits_10m_single ENGINE = MergeTree()
        PARTITION BY toYYYYMM(EventDate)
        ORDER BY (CounterID, EventDate, intHash32(UserID))
        SAMPLE BY intHash32(UserID)
        SETTINGS min_rows_for_wide_part = 0, min_bytes_for_wide_part = 0
    