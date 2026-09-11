-- 特定のタイムスタンプより古い孤立ファイルを削除する
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- ドライラン: 削除対象のファイルをプレビューする
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- データディレクトリのみをスキャンする
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- 位置引数の older_than と名前付き引数を組み合わせる
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
