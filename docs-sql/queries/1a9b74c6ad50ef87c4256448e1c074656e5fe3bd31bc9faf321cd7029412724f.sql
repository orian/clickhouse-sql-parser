-- 位置形式：单个未命名的 older_than 参数
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- 命名形式
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- 无参数：使用所有默认值（older_than = 3 天前）
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
