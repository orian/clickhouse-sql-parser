SELECT h.metric, h.labels, h.histogram, h.count, h.sum
FROM system.metric_log
ARRAY JOIN histograms AS h
WHERE h.metric = 'keeper_response_time_ms' AND h.labels['operation_type'] = 'readonly'
ORDER BY event_time DESC
LIMIT 1;
