
        CREATE TABLE join_dictionary_source_table (key UInt64, value String)
        ENGINE = MergeTree ORDER BY key;
    