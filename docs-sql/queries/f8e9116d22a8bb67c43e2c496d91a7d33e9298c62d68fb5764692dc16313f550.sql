-- Позиционная форма: единственный безымянный аргумент older_than
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- Именованная форма
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- Без аргументов: используются все значения по умолчанию (older_than = 3 дня назад)
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
