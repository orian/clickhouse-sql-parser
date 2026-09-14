
        SELECT dictGetKeys('dict_vector_path', 'value', toUInt32(number % 1000))
        FROM numbers(100000)
        FORMAT NULL;
    