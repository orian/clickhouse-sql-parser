SELECT midpoint(1, toUInt8(3), 0.5) AS result, toTypeName(result) AS type;
-- The type returned is a Float64 as the UInt8 must be promoted to 64 bit for the comparison.
