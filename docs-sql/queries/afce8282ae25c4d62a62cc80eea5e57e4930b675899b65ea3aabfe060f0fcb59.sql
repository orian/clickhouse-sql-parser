/*
 * 두 날짜 사이의 일수를 계산합니다.
 * /* 인수 중 하나라도 NULL이면 NULL을 반환합니다 */
 */
SELECT
    dateDiff('day', toDate('2024-01-01'), toDate('2024-12-31')) AS days_in_year, -- 365
    dateDiff('day', toDate('2020-01-01'), today()) AS days_since  #! since 2020
    ///////////////////////////////////////////////////////////////////
    # TODO: 시간/분 단위 변형 추가
