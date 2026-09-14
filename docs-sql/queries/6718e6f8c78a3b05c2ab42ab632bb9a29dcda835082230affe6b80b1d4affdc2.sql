-- 使用轻量级删除删除 2018 年的所有数据。不建议使用。
DELETE FROM posts WHERE toYear(CreationDate) = 2018
