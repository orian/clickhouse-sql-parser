CREATE TABLE date_interval (id String, start Date, end Date) ENGINE = Memory;
INSERT INTO date_interval VALUES ('a', '2020-01-01', '2020-01-04'), ('a', '2020-01-12', '2020-01-18');

SELECT id, intervalLengthSum(start, end), toTypeName(intervalLengthSum(start, end)) FROM date_interval GROUP BY id ORDER BY id;
