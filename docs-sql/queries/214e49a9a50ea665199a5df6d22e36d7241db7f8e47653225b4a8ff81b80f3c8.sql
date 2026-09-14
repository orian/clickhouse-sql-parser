-- Eliminar archivos huérfanos más antiguos que una marca de tiempo específica
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- Dry run: previsualizar qué archivos se eliminarían
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- Analizar solo el directorio de datos
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- Combinar older_than posicional con argumentos con nombre
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
