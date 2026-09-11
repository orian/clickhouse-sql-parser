SELECT
    toUInt8(257) AS x,
    toTypeName(x),
    reinterpretAsUInt16(x) AS res,
    toTypeName(res)
