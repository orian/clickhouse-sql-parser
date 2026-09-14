
        CREATE TABLE hierarchical_dictionary_source_table
        (
            id UInt64,
            parent_id UInt64
        ) ENGINE = MergeTree ORDER BY id;
    