SELECT kqlBinAt(toDateTime64('2026-08-01 12:34:56', 7, 'UTC'), toIntervalHour(1), toDateTime64('2026-08-01 00:30:00', 7, 'UTC'))
