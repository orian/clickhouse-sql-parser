  -- 온더플라이 뮤테이션을 명시적으로 비활성화합니다
  SET apply_mutations_on_fly = 0;

  SELECT id, v FROM test_on_fly_mutations ORDER BY id;
