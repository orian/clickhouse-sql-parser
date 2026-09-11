SET allow_insert_into_iceberg = 1;

-- Cria alguns snapshots inserindo dados
INSERT INTO iceberg_table VALUES (1);
INSERT INTO iceberg_table VALUES (2);
INSERT INTO iceberg_table VALUES (3);

-- Expira usando apenas a política de retenção
ALTER TABLE iceberg_table EXECUTE expire_snapshots();

-- Expira com fusível de segurança: protege snapshots mais recentes que o timestamp (sintaxe posicional)
ALTER TABLE iceberg_table EXECUTE expire_snapshots('2025-01-01 00:00:00');

-- O mesmo usando a forma de argumento nomeado
ALTER TABLE iceberg_table EXECUTE expire_snapshots(expire_before = '2025-01-01 00:00:00');

-- Substitui os parâmetros de retenção para uma execução
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '3d', retain_last = 10);

-- Expira snapshots explícitos
ALTER TABLE iceberg_table EXECUTE expire_snapshots(snapshot_ids = [101, 102, 103]);

-- Pré-visualização em modo dry-run (sem atualizações de metadados, sem exclusão de arquivos)
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '1d', dry_run = 1);
