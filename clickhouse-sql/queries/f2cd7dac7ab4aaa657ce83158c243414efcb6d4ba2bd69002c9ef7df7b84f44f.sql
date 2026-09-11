
        WITH rand64() % toUInt64({elements_count}) as key
        SELECT dictHas('default.simple_key_hashed_dictionary', key)
        FROM numbers_mt({elements_count})
        FORMAT Null;
    