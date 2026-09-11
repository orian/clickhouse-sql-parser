SELECT
    now() AS current_time,
    now() + toInterval(1, 'hour') AS one_hour_later,
    now() - toInterval(7, 'day') AS week_ago
