SELECT log_time, log_message
FROM logs
WHERE log_time >= toDateTime(1415792726) AND log_time <= toDateTime(1447328726)
