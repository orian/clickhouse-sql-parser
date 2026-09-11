
        WITH rand64() % 5000 as key
        SELECT dictHas('default.simple_key_range_hashed_dictionary', toUInt64(key), key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    