SELECT count(), uniq(pageId) , min(timestamp), max(timestamp) 
FROM session_events 
WHERE type='type3' 
GROUP BY toYear(timestamp) LIMIT 10
