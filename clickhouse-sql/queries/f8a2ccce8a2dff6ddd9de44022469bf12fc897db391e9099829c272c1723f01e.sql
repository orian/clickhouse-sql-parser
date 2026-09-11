
        INSERT INTO cmp128 SELECT
            bitShiftLeft(toUInt128(cityHash64(number, 1)), 64) + cityHash64(number, 2),
            bitShiftLeft(toUInt128(cityHash64(number, 3)), 64) + cityHash64(number, 4),
            toInt128(bitShiftLeft(toUInt128(cityHash64(number, 5)), 64) + cityHash64(number, 6)),
            toInt128(bitShiftLeft(toUInt128(cityHash64(number, 7)), 64) + cityHash64(number, 8)),
            toUInt128(cityHash64(number, 9)),
            toUInt128(cityHash64(number, 10)),
            bitShiftLeft(toUInt128(cityHash64(number, 11)), 64) + cityHash64(number, 12),
            bitShiftLeft(toUInt128(cityHash64(number, if(number % 2 = 0, 11, 13))), 64) + cityHash64(number, 14)
        FROM numbers(10000000)
    