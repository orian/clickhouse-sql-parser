SELECT kqlRangeCount(toDateTime64('2026-08-01 00:00:00', 7, 'UTC'), toDateTime64('2026-08-01 12:00:00', 7, 'UTC'), toIntervalHour(5))
