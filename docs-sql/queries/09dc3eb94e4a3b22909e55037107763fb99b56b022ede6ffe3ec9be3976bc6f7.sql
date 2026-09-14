SELECT avg2(toUInt8(3), 1.0) AS result, toTypeName(result) AS type;
-- UInt8이 비교를 위해 64비트로 승격되어야 하므로 반환되는 유형은 Float64입니다.
