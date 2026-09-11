  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- Desative a materialização em segundo plano das mutações para demonstrar
  -- o comportamento padrão quando as mutações on-the-fly não estão habilitadas
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- Insira algumas linhas na nossa nova tabela
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- Atualize os valores das linhas
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
