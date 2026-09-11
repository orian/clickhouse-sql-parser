
        CREATE TABLE hits_compact AS hits_10m_single ENGINE = MergeTree()
        PARTITION BY toYYYYMM(EventDate)
        ORDER BY (CounterID, EventDate, intHash32(UserID))
        SAMPLE BY intHash32(UserID)
        SETTINGS min_bytes_for_wide_part = '10M'
    