-- This query implies the projection's WHERE, so the projection may be used:
SELECT time, message FROM events WHERE event_type = 'pageview';

-- A stricter query also implies the projection's WHERE:
SELECT time, message FROM events WHERE event_type = 'pageview' AND time > '2024-01-01';

-- This query does NOT imply the projection, so the base table is scanned:
SELECT time, message FROM events WHERE event_type = 'click';
