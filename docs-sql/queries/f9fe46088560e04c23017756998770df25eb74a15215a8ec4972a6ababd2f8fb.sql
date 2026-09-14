-- Следующая дата — 21 апреля 2023 года, это пятница:
SELECT
    toDayOfWeek(toDateTime('2023-04-21')),
    toDayOfWeek(toDateTime('2023-04-21'), 1)
