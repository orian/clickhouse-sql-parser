DETACH TABLE table_name;  -- Required for DROP REPLICA

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- It will remove everything from /table_path_in_zk

ATTACH TABLE table_name;  -- Table will be in readonly mode, because there is no metadata in ZK
