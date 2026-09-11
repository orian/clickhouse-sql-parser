SELECT
    toInt8(-1) AS val,
    toTypeName(val),
    reinterpretAsUInt8(val) AS res,
    toTypeName(res);
