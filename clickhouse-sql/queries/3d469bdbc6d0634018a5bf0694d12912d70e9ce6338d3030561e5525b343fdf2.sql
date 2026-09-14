
        SELECT
            reinterpretAsInt16(a),
            reinterpretAsInt16(b),
            reinterpretAsInt16(c),
            reinterpretAsInt16(d),
            reinterpretAsInt16(f),
            reinterpretAsInt16(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(100000000)
        SETTINGS max_threads = 8
        FORMAT Null
    