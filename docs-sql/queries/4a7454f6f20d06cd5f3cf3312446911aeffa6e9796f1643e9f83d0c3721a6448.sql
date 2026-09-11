SELECT
    toFloat32OrNull('42.7'),
    toFloat32OrNull('NaN'),
    toFloat32OrNull('abc')
FORMAT Vertical
