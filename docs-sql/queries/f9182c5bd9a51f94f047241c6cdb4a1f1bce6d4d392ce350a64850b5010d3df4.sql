SELECT
    toInt128OrNull('-128'),
    toInt128OrNull('abc')
FORMAT Vertical
