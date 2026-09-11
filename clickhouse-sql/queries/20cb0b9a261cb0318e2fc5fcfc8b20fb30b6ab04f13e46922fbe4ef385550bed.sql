
        INSERT INTO simple_key_flat_dictionary_source_table
        SELECT number, number, toString(number), toDecimal64(number, 8), toString(number)
        FROM system.numbers
        LIMIT 5000000;
    