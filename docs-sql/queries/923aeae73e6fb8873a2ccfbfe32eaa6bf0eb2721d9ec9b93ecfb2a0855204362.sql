SELECT
    toInt16(257) AS x,
    toTypeName(x),
    reinterpretAsInt32(x) AS res,
    toTypeName(res)
