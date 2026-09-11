INSERT INTO time_examples VALUES 
    (1, '14:30:25', '14:30:25.123'),
    (2, 52225, 52225.456),
    (3, '26:11:10', '26:11:10.789');  -- 超过 24 小时的值会被规范化

SELECT * FROM time_examples ORDER BY event_id;
