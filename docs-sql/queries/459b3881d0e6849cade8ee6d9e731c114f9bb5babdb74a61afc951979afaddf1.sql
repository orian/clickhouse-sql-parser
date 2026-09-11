-- Parse DateTime
-- - from a string,
-- - from a number interpreted as the number of seconds since 1970-01-01 (a fractional part is truncated to whole seconds).
INSERT INTO dt VALUES ('2019-01-01 00:00:00', 1), (1546300800, 2);

SELECT * FROM dt;
