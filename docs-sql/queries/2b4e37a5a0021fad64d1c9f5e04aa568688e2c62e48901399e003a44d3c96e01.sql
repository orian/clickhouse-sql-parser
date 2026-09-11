DETACH TABLE table_name;  -- DROP REPLICA 的前置操作

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- 将删除 /table_path_in_zk 下的所有内容

ATTACH TABLE table_name;  -- 表将进入只读模式，因为 ZK 中尚无元数据
