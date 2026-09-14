
         SELECT
            reinterpretAsFixedString(a),
            reinterpretAsFixedString(b),
            reinterpretAsFixedString(c),
            reinterpretAsFixedString(d),
            reinterpretAsFixedString(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(50000000)
        SETTINGS max_threads = 8
        FORMAT Null
    