SELECT
    toUInt128(257) AS x,
    toTypeName(x),
    reinterpretAsUInt256(x) AS res,
    toTypeName(res)
