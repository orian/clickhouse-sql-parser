-- 以下日期为 2023 年 4 月 21 日，是星期五：
SELECT
    toDayOfWeek(toDateTime('2023-04-21')),
    toDayOfWeek(toDateTime('2023-04-21'), 1)
