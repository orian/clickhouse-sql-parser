DETACH TABLE table_name;  -- DROP REPLICA に必要

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- /table_path_in_zk 内のすべてを削除します

ATTACH TABLE table_name;  -- ZK にメタデータが存在しないため、テーブルは読み取り専用モードになります
