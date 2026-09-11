SELECT
    toInt64(257) AS x,
    toTypeName(x),
    reinterpretAsInt128(x) AS res,
    toTypeName(res)
