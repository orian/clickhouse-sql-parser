SELECT midpoint(1, toUInt8(3), 0.5) AS result, toTypeName(result) AS type;
-- O tipo retornado é Float64 pois o UInt8 deve ser promovido para 64 bits para a comparação.
