SET allow_insert_into_iceberg = 1;

-- 通过插入数据创建一些快照
INSERT INTO iceberg_table VALUES (1);
INSERT INTO iceberg_table VALUES (2);
INSERT INTO iceberg_table VALUES (3);

-- 仅使用保留策略使快照过期
ALTER TABLE iceberg_table EXECUTE expire_snapshots();

-- 使用安全熔断：保护晚于该时间戳的快照（位置参数语法）
ALTER TABLE iceberg_table EXECUTE expire_snapshots('2025-01-01 00:00:00');

-- 使用命名参数形式实现相同效果
ALTER TABLE iceberg_table EXECUTE expire_snapshots(expire_before = '2025-01-01 00:00:00');

-- 仅针对本次执行覆盖保留参数
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '3d', retain_last = 10);

-- 使指定快照过期
ALTER TABLE iceberg_table EXECUTE expire_snapshots(snapshot_ids = [101, 102, 103]);

-- 试运行预览（不更新元数据，不删除文件）
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '1d', dry_run = 1);
