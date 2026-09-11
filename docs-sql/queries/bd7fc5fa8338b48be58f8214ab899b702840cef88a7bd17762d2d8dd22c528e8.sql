SELECT COUNT(*) 
FROM events 
WHERE event_timestamp > NOW() - INTERVAL '30 days'
  AND event_data->>'additional_info' LIKE 'info_5%';
