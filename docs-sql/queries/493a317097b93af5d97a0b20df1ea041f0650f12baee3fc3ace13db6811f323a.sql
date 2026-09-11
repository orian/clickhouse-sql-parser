INSERT INTO events_raw VALUES
    (now() - INTERVAL 4 SECOND, 101, 'US', 'view', 1),
    (now() - INTERVAL 3 SECOND, 101, 'US', 'click', 1),
    (now() - INTERVAL 2 SECOND, 202, 'DE', 'view', 1),
    (now() - INTERVAL 1 SECOND, 101, 'US', 'view', 1);
