
        WITH (number, toString(number)) as key
        SELECT dictHas('default.complex_key_direct_dictionary', key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    