SELECT midpoint(1, toUInt8(3), 0.5) AS result, toTypeName(result) AS type;
-- UInt8 は比較のために 64 ビットに昇格する必要があるため、返される型は Float64 になります。
