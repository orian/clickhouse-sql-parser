SELECT avg2(toUInt8(3), 1.0) AS result, toTypeName(result) AS type;
-- 返される型はFloat64です。UInt8を比較のために64ビットに昇格させる必要があるためです。
