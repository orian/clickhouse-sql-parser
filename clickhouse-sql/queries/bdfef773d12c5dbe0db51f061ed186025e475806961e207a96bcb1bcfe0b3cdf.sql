
        CREATE TABLE map_subcol_pw_compact (id UInt64, tags Map(String, String))
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = '10G', min_rows_for_wide_part = 1000000000
    