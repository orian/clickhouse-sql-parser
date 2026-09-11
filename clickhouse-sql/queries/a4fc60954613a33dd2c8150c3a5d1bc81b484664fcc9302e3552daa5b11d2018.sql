
        INSERT INTO div128 SELECT
            bitShiftLeft(toUInt128(cityHash64(number, 1)), 64) + cityHash64(number, 2),
            toInt128(bitShiftLeft(toUInt128(cityHash64(number, 3)), 63) + cityHash64(number, 4)),
            toUInt128(cityHash64(number, 5)),
            toInt128(intDiv(cityHash64(number, 6), 2))
        FROM numbers(2000000)
    