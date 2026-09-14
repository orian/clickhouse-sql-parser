-- The result is a sequence of bytes, so it is shown with `hex`.
SELECT
    hex(reinterpretAsString(toDateTime('1970-01-01 01:01:05'))),
    hex(reinterpretAsString(toDate('1970-03-07')))
