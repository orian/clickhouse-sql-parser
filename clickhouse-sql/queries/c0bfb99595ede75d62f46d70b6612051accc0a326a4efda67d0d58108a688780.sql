
        INSERT INTO dict_src_vector_path
        SELECT
            number AS id,
            (number % 10000) AS value,
            concat('value_', toString(number % 10000)) AS value_str
        FROM numbers(10000000);
    