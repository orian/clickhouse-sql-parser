CREATE TABLE qbit (id UInt32, vec QBit(Int8, 2)) ENGINE = Memory;
INSERT INTO qbit VALUES (1, arrayMap(x -> quantizeBFloat16ToInt8(x), [0.1, -0.5]::Array(BFloat16)));
SELECT round(cosineDistanceTransposedQuantized(vec, [0.1, -0.5]::Array(Float32), 8), 6) FROM qbit;
