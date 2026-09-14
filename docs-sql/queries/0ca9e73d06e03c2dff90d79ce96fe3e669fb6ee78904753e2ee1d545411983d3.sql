SET allow_insert_into_iceberg = 1;

-- Crear algunas instantáneas insertando datos
INSERT INTO iceberg_table VALUES (1);
INSERT INTO iceberg_table VALUES (2);
INSERT INTO iceberg_table VALUES (3);

-- Expirar usando solo la política de retención
ALTER TABLE iceberg_table EXECUTE expire_snapshots();

-- Expirar con fusible de seguridad: proteger las instantáneas más recientes que la marca de tiempo (sintaxis posicional)
ALTER TABLE iceberg_table EXECUTE expire_snapshots('2025-01-01 00:00:00');

-- Lo mismo usando la forma de argumento con nombre
ALTER TABLE iceberg_table EXECUTE expire_snapshots(expire_before = '2025-01-01 00:00:00');

-- Sobreescribir los parámetros de retención para una ejecución
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '3d', retain_last = 10);

-- Expirar instantáneas explícitas
ALTER TABLE iceberg_table EXECUTE expire_snapshots(snapshot_ids = [101, 102, 103]);

-- Vista previa en modo simulado (sin actualizaciones de metadatos ni eliminación de archivos)
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '1d', dry_run = 1);
