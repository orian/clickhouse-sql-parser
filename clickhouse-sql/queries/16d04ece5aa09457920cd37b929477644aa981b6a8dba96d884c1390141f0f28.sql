
        SELECT number, concat(toString(number), '-', lowerUTF8(toString(number * 31)), '-', toString(number % 97)) AS bucket
        FROM numbers_mt(10000000)
        ORDER BY number LIMIT 10 BY bucket
        FORMAT NULL
    