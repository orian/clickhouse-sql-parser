-- 删除早于特定时间戳的孤立文件
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- 试运行：预览将被删除的文件
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- 仅扫描数据目录
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- 将位置参数 older_than 与命名参数结合使用
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
