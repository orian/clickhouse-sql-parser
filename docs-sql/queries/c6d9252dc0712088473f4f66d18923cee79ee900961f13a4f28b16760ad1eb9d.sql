SELECT toFixedString('abc', 3) AS a, hex(a), hex(bitShiftRight(a, 12)) AS a_shifted;
