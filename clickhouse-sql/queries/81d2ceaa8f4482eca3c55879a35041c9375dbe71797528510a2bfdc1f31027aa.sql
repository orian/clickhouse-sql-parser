
        CREATE TABLE map_subcol_pw_wide (id UInt64, tags Map(String, String))
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 0, min_rows_for_wide_part = 0
    