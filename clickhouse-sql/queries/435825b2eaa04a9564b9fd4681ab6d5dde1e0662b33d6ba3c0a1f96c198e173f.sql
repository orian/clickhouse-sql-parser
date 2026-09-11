
        SELECT
            reinterpretAsString(a),
            reinterpretAsString(b),
            reinterpretAsString(c),
            reinterpretAsString(d),
            reinterpretAsString(f),
            reinterpretAsString(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(10000000)
        SETTINGS max_threads = 8
        FORMAT Null
    