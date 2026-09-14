
        INSERT INTO hits_mv
        SELECT CounterID, EventDate, UserID, Title
        FROM hits_10m_single
        SETTINGS max_insert_threads=8
    