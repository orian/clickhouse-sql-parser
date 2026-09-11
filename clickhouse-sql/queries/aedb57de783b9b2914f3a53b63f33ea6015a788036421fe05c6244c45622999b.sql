
        INSERT INTO join_dictionary_source_table
        SELECT number, toString(number)
        FROM numbers(1000000);
    