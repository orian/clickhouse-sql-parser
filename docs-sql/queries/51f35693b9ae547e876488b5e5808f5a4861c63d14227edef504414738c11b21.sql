SELECT
    toFloat64OrNull('42.7'),
    toFloat64OrNull('NaN'),
    toFloat64OrNull('abc')
FORMAT Vertical
