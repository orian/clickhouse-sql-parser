SET allow_insert_into_iceberg = 1;

-- Создаём несколько снимков путём вставки данных
INSERT INTO iceberg_table VALUES (1);
INSERT INTO iceberg_table VALUES (2);
INSERT INTO iceberg_table VALUES (3);

-- Удаление устаревших снимков только на основе политики хранения
ALTER TABLE iceberg_table EXECUTE expire_snapshots();

-- Удаление с защитой: снимки новее указанной временной метки не затрагиваются (позиционный синтаксис)
ALTER TABLE iceberg_table EXECUTE expire_snapshots('2025-01-01 00:00:00');

-- То же самое через именованный аргумент
ALTER TABLE iceberg_table EXECUTE expire_snapshots(expire_before = '2025-01-01 00:00:00');

-- Переопределение параметров хранения для одного вызова
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '3d', retain_last = 10);

-- Удаление конкретных снимков по идентификаторам
ALTER TABLE iceberg_table EXECUTE expire_snapshots(snapshot_ids = [101, 102, 103]);

-- Пробный запуск (без обновления метаданных и удаления файлов)
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '1d', dry_run = 1);
