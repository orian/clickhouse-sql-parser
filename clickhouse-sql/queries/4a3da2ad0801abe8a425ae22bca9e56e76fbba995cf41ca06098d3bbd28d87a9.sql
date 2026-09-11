
        SELECT
            reinterpretAsInt64(a),
            reinterpretAsInt64(b),
            reinterpretAsInt64(c),
            reinterpretAsInt64(d),
            reinterpretAsInt64(f),
            reinterpretAsInt64(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(100000000)
        SETTINGS max_threads = 8
        FORMAT Null
    