SELECT toDateTime('2021-04-21 10:20:30', 'America/New_York') AS Time,
toTypeName(Time) AS Type,
timezoneOffset(Time) AS Offset_in_seconds,
(Offset_in_seconds / 3600) AS Offset_in_hours;
