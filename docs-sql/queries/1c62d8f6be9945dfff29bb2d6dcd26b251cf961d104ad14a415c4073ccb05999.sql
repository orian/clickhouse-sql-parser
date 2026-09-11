SELECT
    toInt32(257) AS x,
    toTypeName(x),
    reinterpretAsInt64(x) AS res,
    toTypeName(res)
