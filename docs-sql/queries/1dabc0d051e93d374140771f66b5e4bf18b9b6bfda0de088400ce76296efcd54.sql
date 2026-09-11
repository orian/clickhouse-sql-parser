  CREATE TABLE IF NOT EXISTS float_vs_decimal
  (
  my_float Float64,
  my_decimal Decimal64(3)
  )
  ENGINE=MergeTree
  ORDER BY tuple();

  # 生成 1 000 000 个保留 2 位小数的随机数，并分别以 float 和 decimal 形式存储
  INSERT INTO float_vs_decimal SELECT round(randCanonical(), 3) AS res, res FROM system.numbers LIMIT 1000000;
