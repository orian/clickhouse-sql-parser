
        INSERT INTO FUNCTION file('test_file', '{format}', 'key UInt64, value UInt64')
        SELECT number, number FROM numbers(10000000)
    