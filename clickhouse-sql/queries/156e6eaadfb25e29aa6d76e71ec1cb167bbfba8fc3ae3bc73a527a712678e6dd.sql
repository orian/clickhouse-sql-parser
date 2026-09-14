
        WITH (rand64() % toUInt64(5000) as key, toString(key) as key_id) as complex_key
        SELECT dictGet('default.complex_key_range_hashed_dictionary', 'value', complex_key, key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    