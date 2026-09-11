SET allow_insert_into_iceberg = 1;

-- データを挿入してスナップショットを作成する
INSERT INTO iceberg_table VALUES (1);
INSERT INTO iceberg_table VALUES (2);
INSERT INTO iceberg_table VALUES (3);

-- 保持ポリシーのみを使用してスナップショットを期限切れにする
ALTER TABLE iceberg_table EXECUTE expire_snapshots();

-- 安全ヒューズを使用して期限切れにする: タイムスタンプより新しいスナップショットを保護する（位置引数構文）
ALTER TABLE iceberg_table EXECUTE expire_snapshots('2025-01-01 00:00:00');

-- 名前付き引数形式を使用した場合も同様
ALTER TABLE iceberg_table EXECUTE expire_snapshots(expire_before = '2025-01-01 00:00:00');

-- 1回の実行に限り保持パラメータを上書きする
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '3d', retain_last = 10);

-- 特定のスナップショットを明示的に期限切れにする
ALTER TABLE iceberg_table EXECUTE expire_snapshots(snapshot_ids = [101, 102, 103]);

-- ドライランプレビュー（メタデータの更新なし、ファイルの削除なし）
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '1d', dry_run = 1);
