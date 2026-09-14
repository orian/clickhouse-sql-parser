DETACH TABLE table_name;  -- Necessário para DROP REPLICA

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- Isso removerá tudo de /table_path_in_zk

ATTACH TABLE table_name;  -- A tabela ficará em modo somente leitura, pois não há metadados no ZK
