CREATE TABLE t (value Float64) ENGINE = Memory;
INSERT INTO t VALUES (20.3), (21.1), (21.9), (21.7), (19.9), (21.8);

SELECT studentTTestOneSample()(value, 20.0) FROM t;
