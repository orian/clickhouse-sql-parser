EXPLAIN PIPELINE graph = 1, compact = 0 
SELECT count(), uniq(pageId) , min(timestamp), max(timestamp) 
FROM session_events 
WHERE type='type3' 
GROUP BY toYear(timestamp) 
LIMIT 10 
SETTINGS allow_experimental_parallel_reading_from_replicas=0 
FORMAT TSV;
