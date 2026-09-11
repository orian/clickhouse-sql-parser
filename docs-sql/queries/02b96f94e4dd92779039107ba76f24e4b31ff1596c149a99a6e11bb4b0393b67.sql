SELECT
    toInt64OrNull('-64'),
    toInt64OrNull('abc')
FORMAT Vertical
