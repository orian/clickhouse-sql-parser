  CREATE TABLE IF NOT EXISTS float_vs_decimal
  (
  my_float Float64,
  my_decimal Decimal64(3)
  )
  ENGINE=MergeTree
  ORDER BY tuple();

  # Generar 1 000 000 números aleatorios con 2 posición decimal y almacenarlos como float y como decimal
  INSERT INTO float_vs_decimal SELECT round(randCanonical(), 3) AS res, res FROM system.numbers LIMIT 1000000;
