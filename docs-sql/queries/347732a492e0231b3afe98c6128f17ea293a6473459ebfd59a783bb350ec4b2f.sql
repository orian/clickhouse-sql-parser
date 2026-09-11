-- (1) 暂停写入
-- (2) 创建目标表的副本

CREATE TABLE pypi_v2 AS pypi

SELECT count() FROM pypi_v2
