  CREATE TABLE IF NOT EXISTS float_vs_decimal
  (
  my_float Float64,
  my_decimal Decimal64(3)
  )
  ENGINE=MergeTree
  ORDER BY tuple();

  # 소수점 이하 2자리 난수 1,000,000개를 생성하여 float와 decimal로 저장합니다
  INSERT INTO float_vs_decimal SELECT round(randCanonical(), 3) AS res, res FROM system.numbers LIMIT 1000000;
