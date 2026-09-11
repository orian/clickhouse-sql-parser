WITH toDateTime64('2025-02-13 23:31:31.011123456', 9, 'UTC') AS dt64
SELECT toUnixTimestamp64Nano(dt64);
