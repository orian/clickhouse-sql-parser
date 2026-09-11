-- Q3.1: 주말 동안 실내 온도가 영하로 내려갔습니까?

SELECT *
FROM logs3
WHERE event_type = 'temperature'
  AND event_value <= 32.0
  AND log_time >= '2019-11-29 17:00:00.000';
