SYSTEM RESTORE REPLICA table_name;  -- Irá desanexar todas as partições, recriar os metadados no ZK (como se fosse uma nova tabela vazia) e então reanexar todas as partições

SYSTEM SYNC REPLICA table_name; -- Aguarda as réplicas sincronizarem as partes. Também é recomendável verificar `system.detached_parts` em todas as réplicas após a conclusão da recuperação.
