
        WITH rand64() % toUInt64({elements_count}) as key
        SELECT dictHas('default.simple_key_direct_dictionary', key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    