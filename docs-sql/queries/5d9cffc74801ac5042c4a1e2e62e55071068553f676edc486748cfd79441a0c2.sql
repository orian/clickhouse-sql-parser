  CREATE TABLE IF NOT EXISTS float_vs_decimal
  (
  my_float Float64,
  my_decimal Decimal64(3)
  )
  ENGINE=MergeTree
  ORDER BY tuple();

  # Génère 1 000 000 nombres aléatoires avec 2 décimales et les stocke sous forme de float et de decimal
  INSERT INTO float_vs_decimal SELECT round(randCanonical(), 3) AS res, res FROM system.numbers LIMIT 1000000;
