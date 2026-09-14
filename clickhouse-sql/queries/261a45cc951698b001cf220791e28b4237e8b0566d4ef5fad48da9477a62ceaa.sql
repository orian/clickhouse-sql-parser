
        SELECT
            reinterpretAsInt256(a),
            reinterpretAsInt256(b),
            reinterpretAsInt256(c),
            reinterpretAsInt256(d),
            reinterpretAsInt256(f),
            reinterpretAsInt256(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(5000000)
        SETTINGS max_threads = 8
        FORMAT Null
    