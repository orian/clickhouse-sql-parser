SELECT u.country, COUNT(*) as events, AVG(LENGTH(e.event_data::text))::int as avg_json_size
FROM events e
JOIN users u ON e.user_id = u.user_id
GROUP BY u.country
ORDER BY events DESC;
