-- Parse DateTime64
-- - from an integer interpreted as the number of seconds since 1970-01-01 (like DateTime),
-- - from a decimal interpreted as the number of seconds, the fractional part giving sub-second precision,
-- - from a string.

INSERT INTO dt64
VALUES
(1546300800, 1),
(1546300800.123, 2),
('2019-01-01 00:00:00', 3);

SELECT * FROM dt64;
