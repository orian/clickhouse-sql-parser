  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- Отключим фоновую материализацию мутаций, чтобы показать
  -- поведение по умолчанию, когда мутации на лету не включены
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- Вставим несколько строк в нашу новую таблицу
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- Обновим значения в строках
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
