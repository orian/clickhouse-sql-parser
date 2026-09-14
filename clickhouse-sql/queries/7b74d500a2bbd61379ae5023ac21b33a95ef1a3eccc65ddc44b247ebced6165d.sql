
        CREATE TABLE volume_reducing_functions (key UInt64, s String, fs FixedString(64))
        ENGINE = MergeTree ORDER BY tuple()
    