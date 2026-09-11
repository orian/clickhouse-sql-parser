SELECT toFixedString('abc', 3) AS a, hex(a), hex(bitShiftLeft(a, 4)) AS a_shifted;
