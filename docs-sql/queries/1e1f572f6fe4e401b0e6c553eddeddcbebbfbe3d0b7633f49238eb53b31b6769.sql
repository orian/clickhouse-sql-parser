SELECT toDateTime64(1546300800.000, 3) AS value, toTypeName(value);
-- Without the decimal point the value is still treated as Unix Timestamp in seconds
SELECT toDateTime64(1546300800000, 3) AS value, toTypeName(value);
