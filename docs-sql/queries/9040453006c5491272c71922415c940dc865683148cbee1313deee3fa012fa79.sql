-- The shifted value is binary, so it is shown with `hex`.
SELECT 'abc' AS a, hex(a), hex(bitShiftRight(a, 12)) AS a_shifted;
