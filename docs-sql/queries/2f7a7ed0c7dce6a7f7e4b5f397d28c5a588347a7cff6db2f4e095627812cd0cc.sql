SELECT midpoint(1, toUInt8(3), 0.5) AS result, toTypeName(result) AS type;
-- Возвращаемый тип — Float64, так как UInt8 необходимо расширить до 64 бит для выполнения сравнения.
