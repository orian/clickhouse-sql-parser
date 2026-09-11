-- Mapの特定のキーをデフォルト値とするカラムを追加することで解決する
ALTER TABLE tbl ADD COLUMN hostname LowCardinality(String) DEFAULT ResourceAttributes['1'];
ALTER TABLE tbl MATERIALIZE COLUMN hostname;

-- hostname（新しいカラム）へのクエリが高速化された
-- [cold] 0 rows in set. Elapsed: 2.215 sec. Processed 1.00 million rows, 21.67 MB (451.52 thousand rows/s., 9.78 MB/s.)
-- [warm] 0 rows in set. Elapsed: 0.541 sec. Processed 1.00 million rows, 21.67 MB (1.85 million rows/s., 40.04 MB/s.)
SELECT
    COUNT(*),
    avg(Duration/1E6) as average,
    quantile(0.95)(Duration/1E6) as p95,
    quantile(0.99)(Duration/1E6) as p99,
    hostname
FROM tbl
GROUP BY hostname ORDER BY 1 DESC LIMIT 50 FORMAT Null;

-- キャッシュをクリアしてコールドクエリを実行する
SYSTEM DROP FILESYSTEM CACHE;
