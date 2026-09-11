SELECT
    toUInt16(257) AS x,
    toTypeName(x),
    reinterpretAsUInt32(x) AS res,
    toTypeName(res)
