DETACH TABLE table_name;  -- Требуется для DROP REPLICA

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- Удалит всё из /table_path_in_zk

ATTACH TABLE table_name;  -- Таблица будет в режиме только для чтения, поскольку метаданные в ZK отсутствуют
