-- 位置指定形式: 名前なしの older_than 引数を1つ指定
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- 名前付き形式
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- 引数なし: すべてデフォルト値を使用 (older_than = 3日前)
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
