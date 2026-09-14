  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- Deshabilita la materialización en segundo plano de las mutaciones para mostrar
  -- el comportamiento predeterminado cuando las mutaciones sobre la marcha no están habilitadas
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- Inserta algunas filas en la nueva tabla
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- Actualiza los valores de las filas
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
