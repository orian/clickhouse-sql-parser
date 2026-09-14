
        SELECT dictGetKeys('dict_vector_path', 'value_str', concat('value_', toString(number % 1000)))
        FROM numbers(100000)
        FORMAT NULL;
    