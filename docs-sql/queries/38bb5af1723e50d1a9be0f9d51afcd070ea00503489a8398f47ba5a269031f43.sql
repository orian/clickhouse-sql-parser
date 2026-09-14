CREATE TABLE dt_interval (id String, start DateTime, end DateTime) ENGINE = Memory;
INSERT INTO dt_interval VALUES ('a', '2020-01-01 01:12:30', '2020-01-01 02:10:10'), ('a', '2020-01-01 02:05:30', '2020-01-01 02:50:31'), ('a', '2020-01-01 03:11:22', '2020-01-01 03:23:31');

SELECT id, intervalLengthSum(start, end), toTypeName(intervalLengthSum(start, end)) FROM dt_interval GROUP BY id ORDER BY id;
