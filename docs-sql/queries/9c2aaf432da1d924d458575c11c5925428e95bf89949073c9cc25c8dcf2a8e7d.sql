SELECT avg2(toUInt8(3), 1.0) AS result, toTypeName(result) AS type;
-- 返回类型为 Float64，因为 UInt8 必须提升为 64 位才能进行比较。
