-- 다음 날짜는 2023년 4월 21일이며, 금요일이었습니다:
SELECT
    toDayOfWeek(toDateTime('2023-04-21')),
    toDayOfWeek(toDateTime('2023-04-21'), 1)
