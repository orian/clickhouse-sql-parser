-- The padding is made of null bytes, so the result is shown with `hex`.
SELECT hex(toFixedString('foo', 8)) AS s;
