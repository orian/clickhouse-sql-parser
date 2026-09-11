
        INSERT INTO test_group_by_strings
        SELECT
            lpad(toString(number % 1000), 8, '0'),
            lpad(toString(((number * 17 + 12345) % 1000)), 8, '0'),
            lpad(toString(number % 1000), 16, '0'),
            lpad(toString(((number * 17 + 12345) % 1000)), 16, '0'),
            lpad(toString(number % 1000), 32, '0'),
            lpad(toString(((number * 17 + 12345) % 1000)), 32, '0'),
            lpad(toString(number % 1000), 64, '0'),
            lpad(toString(((number * 17 + 12345) % 1000)), 64, '0'),
            lpad(toString(number % 1000), 512, '0'),
            lpad(toString(((number * 17 + 12345) % 1000)), 512, '0'),
            number
        FROM numbers_mt(3000000)
    