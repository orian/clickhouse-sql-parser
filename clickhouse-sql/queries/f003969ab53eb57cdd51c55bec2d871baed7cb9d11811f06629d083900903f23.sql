
        CREATE TABLE map_subcol_pw_mixed (id UInt64, tags Map(String, String), value UInt32)
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 0, min_rows_for_wide_part = 0
    