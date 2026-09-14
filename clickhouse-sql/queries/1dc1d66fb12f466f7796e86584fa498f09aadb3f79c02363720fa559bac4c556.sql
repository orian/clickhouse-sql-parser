
        WITH rand64() % toUInt64({elements_count}) as key
        SELECT dictGet('default.simple_key_direct_dictionary', ('value_int', 'value_string', 'value_decimal', 'value_string_nullable'), key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    