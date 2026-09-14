
        SELECT wide
        FROM remote('127.0.0.{{2,3}}',
             view(SELECT repeat(toString(number), 1000) AS wide, range(500) AS arr FROM numbers(500)))
        ARRAY JOIN arr
        FORMAT Null
    