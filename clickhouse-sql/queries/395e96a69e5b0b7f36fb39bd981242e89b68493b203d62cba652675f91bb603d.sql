
        WITH rand64() % toUInt64({elements_count}) as key
        SELECT dictGet('default.simple_key_direct_dictionary', {column_name}, key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    