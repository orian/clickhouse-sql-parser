SET enable_time_time64_type = 1;
SET use_legacy_to_time = 0;
SELECT toTime(toDateTime(52225, 'UTC'))
