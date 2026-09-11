    SET param_date = '2024-01-15';
    SET param_timestamp = '2024-01-15 14:30:00';

    SELECT * FROM events
    WHERE event_date = {date: Date}
       OR event_timestamp > {timestamp: DateTime};
