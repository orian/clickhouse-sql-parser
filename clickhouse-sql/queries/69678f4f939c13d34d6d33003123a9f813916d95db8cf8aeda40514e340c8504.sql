
        INSERT INTO simple_key_range_hashed_dictionary_source_table
        SELECT key, key, range_start * 2, range_start * 2 + 1 FROM
            (SELECT number as key FROM numbers(10000)) as keys,
            (SELECT number as range_start FROM numbers(1000)) as ranges;
    