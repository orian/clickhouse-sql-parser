  -- Habilite as mutações on-the-fly
  SET apply_mutations_on_fly = 1;

  SELECT id, v FROM test_on_fly_mutations ORDER BY id;
