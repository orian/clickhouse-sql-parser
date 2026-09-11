CREATE TABLE t (response_time UInt32, weight UInt32) ENGINE = Memory;
INSERT INTO t VALUES (68, 1), (104, 2), (112, 3), (126, 2), (138, 1), (162, 1);

SELECT quantilesTimingWeighted(0.5, 0.99)(response_time, weight) FROM t;
