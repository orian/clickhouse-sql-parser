  -- compact パーツを持つテーブルを作成
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- min_bytes_for_wide_part = 10485760 のデフォルト値を超えるには十分なサイズではない

  -- パーツのタイプを確認
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- compact テーブルの圧縮後および非圧縮のカラムサイズを取得
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- wide パーツを持つテーブルを作成 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- パーツのタイプを確認
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- wide テーブルの圧縮後および非圧縮サイズを取得
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
