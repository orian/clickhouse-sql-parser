-- (1) Приостановить вставки
-- (2) Создать дубликат целевой таблицы

CREATE TABLE pypi_v2 AS pypi

SELECT count() FROM pypi_v2
