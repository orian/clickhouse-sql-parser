-- 以下の日付は2023年4月21日（金曜日）です:
SELECT
    toDayOfWeek(toDateTime('2023-04-21')),
    toDayOfWeek(toDateTime('2023-04-21'), 1)
