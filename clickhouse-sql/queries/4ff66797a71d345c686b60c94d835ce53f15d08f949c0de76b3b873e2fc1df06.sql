
        INSERT INTO complex_key_direct_dictionary_source_table
        SELECT number, toString(number), number, toString(number), toDecimal64(number, 8), toString(number)
        FROM system.numbers
        LIMIT 50000;
    