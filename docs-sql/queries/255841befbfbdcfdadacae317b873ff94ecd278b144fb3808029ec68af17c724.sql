CREATE TABLE prod
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

CREATE TABLE staging
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

-- Начальные данные
INSERT INTO prod VALUES (123, 'John', 33);
INSERT INTO prod VALUES (456, 'Ksenia', 48);
-- Загрузка данных
INSERT INTO staging VALUES (8811, 'Alice', 50);
INSERT INTO staging VALUES (8812, 'Bob', 23);

-- Проверка импорта
SELECT 'Staging count:', COUNT() FROM staging;
-- Перемещение партиции
ALTER TABLE staging MOVE PARTITION tuple() TO TABLE prod; -- атомарная операция

-- Проверка данных
SELECT 'Prod count:', COUNT() FROM prod;
SELECT * FROM prod;
