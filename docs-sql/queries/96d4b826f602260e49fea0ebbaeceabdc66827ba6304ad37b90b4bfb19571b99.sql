  -- تعطيل on-the-fly mutations صراحةً
  SET apply_mutations_on_fly = 0;

  SELECT id, v FROM test_on_fly_mutations ORDER BY id;
