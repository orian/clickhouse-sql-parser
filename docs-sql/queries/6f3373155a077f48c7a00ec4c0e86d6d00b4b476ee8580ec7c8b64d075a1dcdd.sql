-- FINAL을 사용하여 올바른 현재 상태(state) 조회
SELECT * FROM your_table_name FINAL WHERE __DELETED = 'false';
SELECT * FROM your_table_name FINAL LIMIT 10;
SELECT * FROM your_table_name FINAL WHERE <filter by keys in ORDER BY clause>;
SELECT count(*) FROM your_table_name FINAL;
