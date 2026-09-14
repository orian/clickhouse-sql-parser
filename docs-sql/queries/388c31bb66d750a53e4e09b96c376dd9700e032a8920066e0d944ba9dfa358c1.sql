SYSTEM RESTORE REPLICA table_name;  -- 将分离所有分区，在 ZK 中重新创建元数据（如同新建空表），然后重新挂载所有分区

SYSTEM SYNC REPLICA table_name; -- 等待副本同步 parts。建议在恢复完成后检查所有副本上的 `system.detached_parts`。
