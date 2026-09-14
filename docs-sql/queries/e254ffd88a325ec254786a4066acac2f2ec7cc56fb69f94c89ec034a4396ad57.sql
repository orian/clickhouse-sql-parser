    INSERT INTO events (event_name, event_type, event_timestamp, event_data, user_id, user_ip)
    SELECT
       'Event ' || gs::text AS event_name,
       CASE
           WHEN random() < 0.5 THEN 'click'
           WHEN random() < 0.75 THEN 'view'
           WHEN random() < 0.9 THEN 'purchase'
           WHEN random() < 0.98 THEN 'signup'
           ELSE 'logout'
       END AS event_type,
       NOW() - INTERVAL '1 day' * (gs % 365) AS event_timestamp,
       jsonb_build_object('key', 'value' || gs::text, 'additional_info', 'info_' || (gs % 100)::text) AS event_data,
       GREATEST(1, LEAST(1000, FLOOR(POWER(random(), 2) * 1000) + 1)) AS user_id,
       ('192.168.1.' || ((gs % 254) + 1))::inet AS user_ip
    FROM
       generate_series(1, 1000000) gs;
