  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- 예시를 위해 뮤테이션의 백그라운드 머티리얼라이즈를 비활성화합니다
  -- 온더플라이 뮤테이션이 활성화되지 않았을 때의 기본 동작을 보여줍니다
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- 새 테이블에 몇 개의 행을 삽입합니다
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- 행 값을 업데이트합니다
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
