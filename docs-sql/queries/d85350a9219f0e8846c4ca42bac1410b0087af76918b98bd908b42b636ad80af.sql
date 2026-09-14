  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- Désactiver la matérialisation des mutations en arrière-plan pour illustrer
  -- le comportement par défaut lorsque les mutations à la volée ne sont pas activées
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- Insérer quelques lignes dans notre nouvelle table
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- Mettre à jour les valeurs des lignes
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
