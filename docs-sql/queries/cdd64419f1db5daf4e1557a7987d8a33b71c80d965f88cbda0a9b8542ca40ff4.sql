SELECT midpoint(1, toUInt8(3), 0.5) AS result, toTypeName(result) AS type;
-- 반환 유형은 Float64입니다. UInt8이 비교를 위해 64비트로 승격되어야 하기 때문입니다.
