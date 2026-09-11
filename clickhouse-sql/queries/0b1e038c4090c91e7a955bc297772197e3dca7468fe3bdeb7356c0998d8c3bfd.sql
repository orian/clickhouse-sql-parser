
        SELECT COUNT() 
        FROM join_dictionary_source_table
        JOIN join_hashed_dictionary
        ON join_dictionary_source_table.key = toUInt64(join_hashed_dictionary.key);
    