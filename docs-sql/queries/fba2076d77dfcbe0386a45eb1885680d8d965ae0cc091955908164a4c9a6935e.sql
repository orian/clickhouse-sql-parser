SET enable_time_time64_type = 1;
SET use_legacy_to_time = 0;
SELECT toTime(toDateTime64('2025-04-15 14:30:25.123', 3))
