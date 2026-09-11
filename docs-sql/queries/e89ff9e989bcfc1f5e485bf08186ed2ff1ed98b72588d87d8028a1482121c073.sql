-- The result is a sequence of bytes, so it is shown with `hex`.
SELECT
    hex(reinterpretAsFixedString(toDateTime('1970-01-01 01:01:05'))),
    hex(reinterpretAsFixedString(toDate('1970-03-07')))
