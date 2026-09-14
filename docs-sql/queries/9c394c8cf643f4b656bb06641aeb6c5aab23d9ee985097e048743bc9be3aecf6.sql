-- Forma posicional: único argumento older_than sem nome
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- Forma nomeada
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- Sem argumentos: usa todos os padrões (older_than = 3 dias atrás)
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
