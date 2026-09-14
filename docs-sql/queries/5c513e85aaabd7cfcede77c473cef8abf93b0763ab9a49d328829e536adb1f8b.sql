CREATE TABLE t (response_time UInt32) ENGINE = Memory;
INSERT INTO t VALUES (72), (112), (126), (145), (104), (242), (313), (168), (108);

SELECT quantileTiming(response_time) FROM t;
