-- The shifted value is binary, so it is shown with `hex`.
SELECT 'abc' AS a, hex(a), hex(bitShiftLeft(a, 4)) AS a_shifted;
