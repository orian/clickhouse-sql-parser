-- 首先，确认用于存储数据的数据库。
SELECT * FROM [database].connect_state

-- 找到与该 topic 和分区匹配的键。
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
