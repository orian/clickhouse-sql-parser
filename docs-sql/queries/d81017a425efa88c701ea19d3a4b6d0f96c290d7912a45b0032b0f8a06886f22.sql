SET use_legacy_to_time = 1;
SELECT toTime(toDateTime64('1970-12-10 01:20:30.3000', 3)) AS result, toTypeName(result)
