
        INSERT INTO FUNCTION file('test_data_sparse.json', LineAsString)
        SELECT '{{"id": ' || number || ', "c' || number % 50 || '": "' || hex(rand()) || '"}}'
        FROM numbers(100000) SETTINGS engine_file_truncate_on_insert = 1
    