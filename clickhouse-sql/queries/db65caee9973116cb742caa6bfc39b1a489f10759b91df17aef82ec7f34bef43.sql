
        SELECT
            reinterpretAsUInt16(a),
            reinterpretAsUInt16(b),
            reinterpretAsUInt16(c),
            reinterpretAsUInt16(d),
            reinterpretAsUInt16(f),
            reinterpretAsUInt16(g),

            toUInt64(number) as a,
            toUInt256(number) as b,
            toInt128(number) as c,
            toInt256(number) as d,
            toString(number) as f,
            toFixedString(f, 20) as g
        FROM numbers_mt(100000000)
        SETTINGS max_threads = 8
        FORMAT Null
    