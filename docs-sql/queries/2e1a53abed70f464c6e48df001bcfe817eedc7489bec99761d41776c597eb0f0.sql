-- The type returned is a Float64 as the UInt8 must be promoted to 64 bit for the comparison.
SELECT least(1, 2, toUInt8(3), 3.) AS result, toTypeName(result) AS type;
