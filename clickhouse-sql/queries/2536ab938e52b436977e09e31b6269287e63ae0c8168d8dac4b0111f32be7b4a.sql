
        INSERT INTO FUNCTION file('test_file', '{format}', 'key UInt64, value1 UInt64, value2 UInt64, value3 UInt64, value4 UInt64, value5 UInt64')
        SELECT number, number, number, number, number, number FROM numbers(1000000)
    