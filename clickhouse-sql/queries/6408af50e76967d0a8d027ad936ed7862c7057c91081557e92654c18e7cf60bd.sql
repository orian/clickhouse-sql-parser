
        WITH (rand64() % toUInt64({elements_count}), toString(rand64() % toUInt64({elements_count}))) as key
        SELECT dictGet('default.complex_key_hashed_array_dictionary', {column_name}, key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    