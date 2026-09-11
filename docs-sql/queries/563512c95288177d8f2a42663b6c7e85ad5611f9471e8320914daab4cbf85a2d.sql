CREATE TABLE qbit (id UInt32, vec QBit(Float64, 2)) ENGINE = Memory;
INSERT INTO qbit VALUES (1, [0, 1]);
SELECT L2DistanceTransposed(vec, array(1, 2), 16) FROM qbit;
