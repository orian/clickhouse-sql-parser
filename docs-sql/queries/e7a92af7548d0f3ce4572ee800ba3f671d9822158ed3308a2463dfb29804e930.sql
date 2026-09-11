-- 지난 1시간 동안 가장 활동적인 사용자 상위 10명 조회
SELECT user\_id, COUNT(*) AS actionsFROM analytics.eventsWHERE event\_time >= now() - INTERVAL 1 HOURGROUP BY user\_idORDER BY actions DESCLIMIT 10;
