SYSTEM RESTORE REPLICA table_name;  -- Desconectará todas las particiones, recreará los metadatos en ZK (como si fuera una tabla nueva vacía) y luego volverá a adjuntar todas las particiones

SYSTEM SYNC REPLICA table_name; -- Espera a que las réplicas sincronicen las partes. También se recomienda verificar `system.detached_parts` en todas las réplicas una vez finalizada la recuperación.
