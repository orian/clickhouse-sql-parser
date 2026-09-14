-- 使用 FINAL 获取正确的当前状态
SELECT * FROM your_table_name FINAL WHERE __DELETED = 'false';
SELECT * FROM your_table_name FINAL LIMIT 10;
SELECT * FROM your_table_name FINAL WHERE <filter by keys in ORDER BY clause>;
SELECT count(*) FROM your_table_name FINAL;
