SELECT * FROM table WHERE user_id = 'alice';
SELECT * FROM table WHERE user_id IN ('alice', 'bob');
SELECT * FROM table WHERE has(['alice', 'bob'], user_id);
