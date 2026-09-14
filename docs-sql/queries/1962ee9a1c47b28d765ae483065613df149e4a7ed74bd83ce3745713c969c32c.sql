-- Forma posicional: argumento older_than sin nombre
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- Forma con nombre
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- Sin argumentos: usar todos los valores predeterminados (older_than = hace 3 días)
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
