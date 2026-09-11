SELECT _path, count() AS rows
FROM my_lake.`logs.application`
WHERE event_time >= '2026-03-01'
  AND event_time < '2026-03-02'
GROUP BY _path
ORDER BY rows DESC;
