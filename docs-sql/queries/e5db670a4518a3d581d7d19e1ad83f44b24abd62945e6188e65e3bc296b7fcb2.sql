-- Q3.1: 週末に室内温度が氷点下に達しましたか？

SELECT *
FROM logs3
WHERE event_type = 'temperature'
  AND event_value <= 32.0
  AND log_time >= '2019-11-29 17:00:00.000';
