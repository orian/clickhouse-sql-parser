    -- Via ClickHouse
    SELECT 
        u.country,
        u.platform,
        COUNT(DISTINCT e.user_id) as users,
        COUNT(*) as total_events,
        ROUND(COUNT(*)::numeric / COUNT(DISTINCT e.user_id), 2) as events_per_user,
        SUM(CASE WHEN e.event_type = 'purchase' THEN 1 ELSE 0 END) as purchases
    FROM organization.public_events e
    JOIN organization.public_users u ON e.user_id = u.user_id
    GROUP BY u.country, u.platform
    ORDER BY total_events DESC
    LIMIT 10;
