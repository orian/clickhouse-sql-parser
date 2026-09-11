SET enable_time_time64_type = 1;
SELECT toTime64(toDateTime64('2025-04-15 14:30:25.123456', 6), 6)
