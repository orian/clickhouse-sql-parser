    SELECT event_type, COUNT(*) as count 
    FROM events 
    GROUP BY event_type 
    ORDER BY count DESC;
