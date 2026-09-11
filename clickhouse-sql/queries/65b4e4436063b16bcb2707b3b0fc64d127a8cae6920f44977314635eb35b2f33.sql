
        INSERT INTO div256 SELECT
            bitShiftLeft(toUInt256(cityHash64(number, 1)), 192) + bitShiftLeft(toUInt256(cityHash64(number, 2)), 128)
                + bitShiftLeft(toUInt256(cityHash64(number, 3)), 64) + cityHash64(number, 4),
            toInt256(bitShiftLeft(toUInt256(cityHash64(number, 5)), 191) + bitShiftLeft(toUInt256(cityHash64(number, 6)), 128)
                + bitShiftLeft(toUInt256(cityHash64(number, 7)), 64) + cityHash64(number, 8)),
            toUInt256(cityHash64(number, 9)),
            toDecimal256(number, 40) / 7
        FROM numbers(2000000)
    