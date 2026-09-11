-- Remove arquivos órfãos mais antigos que um timestamp específico
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- Dry run: visualize quais arquivos seriam excluídos
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- Escanear apenas o diretório de dados
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- Combinar older_than posicional com argumentos nomeados
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
