
        CREATE TABLE vrow_filtered (a UInt64, b UInt8)
        ENGINE = MergeTree ORDER BY a
        SETTINGS parts_to_throw_insert = 100000, parts_to_delay_insert = 100000
    