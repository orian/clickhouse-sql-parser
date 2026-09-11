WITH toDateTime64('2025-02-13 23:31:31.011123', 6, 'UTC') AS dt64
SELECT toUnixTimestamp64Micro(dt64);
