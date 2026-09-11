
        WITH rand64() % 5000 as key
        SELECT dictGet('default.simple_key_range_hashed_dictionary', 'value', toUInt64(key), key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    