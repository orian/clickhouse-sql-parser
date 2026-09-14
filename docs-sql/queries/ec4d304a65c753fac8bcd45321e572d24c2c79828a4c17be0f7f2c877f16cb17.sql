    -- Via ClickHouse
    SELECT 
        user_id,
        COUNT(*) as total_events,
        COUNT(DISTINCT event_type) as unique_event_types,
        SUM(CASE WHEN event_type = 'purchase' THEN 1 ELSE 0 END) as purchases,
        MIN(event_timestamp) as first_event,
        MAX(event_timestamp) as last_event
    FROM organization.public_events
    GROUP BY user_id
    ORDER BY total_events DESC
    LIMIT 10;
