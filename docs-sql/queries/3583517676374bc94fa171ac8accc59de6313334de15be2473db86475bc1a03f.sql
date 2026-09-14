  CREATE TABLE IF NOT EXISTS float_vs_decimal
  (
  my_float Float64,
  my_decimal Decimal64(3)
  )
  ENGINE=MergeTree
  ORDER BY tuple();

  # 小数点以下 2 桁のランダムな数値を 1 000 000 個生成し、float と decimal として保存する
  INSERT INTO float_vs_decimal SELECT round(randCanonical(), 3) AS res, res FROM system.numbers LIMIT 1000000;
