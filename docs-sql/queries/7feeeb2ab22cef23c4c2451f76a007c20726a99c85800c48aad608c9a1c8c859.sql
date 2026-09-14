  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- 挙動を確認するため、ミューテーションのバックグラウンドでのマテリアライゼーションを無効にする
  -- オンザフライミューテーションが有効でない場合のデフォルトの挙動を示す
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- 新しいテーブルにいくつかの行を挿入する
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- 行の値を更新する
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
