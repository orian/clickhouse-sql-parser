
        INSERT INTO minmax_wide SELECT
            bitShiftLeft(toUInt128(cityHash64(number, 1)), 64) + cityHash64(number, 2),
            toInt128(bitShiftLeft(toUInt128(cityHash64(number, 3)), 64) + cityHash64(number, 4)),
            toInt128(cityHash64(number, 5)),
            toInt128(number),
            toDecimal128(cityHash64(number, 6) % 1000000000, 4),
            toDecimal256(cityHash64(number, 7) % 1000000000, 8),
            toInt256(bitShiftLeft(toUInt128(cityHash64(number, 8)), 64) + cityHash64(number, 9)),
            number % 7 = 0,
            if(number % 5 = 0, NULL, toInt128(cityHash64(number, 10)))
        FROM numbers(10000000)
    