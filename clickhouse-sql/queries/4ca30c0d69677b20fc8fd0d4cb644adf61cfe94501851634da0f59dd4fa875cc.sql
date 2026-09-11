
        SELECT
            reinterpretAsInt128(a),
            reinterpretAsInt128(b),
            reinterpretAsInt128(c),
            reinterpretAsInt128(d),
            reinterpretAsInt128(f),
            reinterpretAsInt128(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(50000000)
        SETTINGS max_threads = 8
        FORMAT Null
    