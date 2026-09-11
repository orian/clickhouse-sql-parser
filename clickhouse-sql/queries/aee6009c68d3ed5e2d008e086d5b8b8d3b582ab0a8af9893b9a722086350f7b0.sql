
        WITH (number, toString(number)) as key
        SELECT dictGet('default.complex_key_direct_dictionary', {column_name}, key)
        FROM system.numbers
        LIMIT {elements_count}
        FORMAT Null;
    