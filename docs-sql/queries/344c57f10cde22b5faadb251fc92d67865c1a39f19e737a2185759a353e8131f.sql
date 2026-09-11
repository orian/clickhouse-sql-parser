SELECT toDateTime('2019-01-01 00:00:00', 'UTC') AS time_utc,
toTypeName(time_utc) AS type_utc,
toInt32(time_utc) AS int32utc,
toTimezone(time_utc, 'Asia/Yekaterinburg') AS time_yekat,
toTypeName(time_yekat) AS type_yekat,
toInt32(time_yekat) AS int32yekat,
toTimezone(time_utc, 'US/Samoa') AS time_samoa,
toTypeName(time_samoa) AS type_samoa,
toInt32(time_samoa) AS int32samoa
FORMAT Vertical;
