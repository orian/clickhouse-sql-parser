DETACH TABLE table_name;  -- Necesario para DROP REPLICA

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- Eliminará todo lo que haya en /table_path_in_zk

ATTACH TABLE table_name;  -- La tabla estará en modo de solo lectura, ya que no hay metadatos en ZK
