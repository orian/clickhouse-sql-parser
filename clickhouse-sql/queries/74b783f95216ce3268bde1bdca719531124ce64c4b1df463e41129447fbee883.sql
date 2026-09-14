
        WITH rand64() % toUInt64({elements_count}) as key
        SELECT dictGet('default.simple_key_hashed_dictionary', {column_name}, key)
        FROM numbers_mt({elements_count})
        FORMAT Null;
    