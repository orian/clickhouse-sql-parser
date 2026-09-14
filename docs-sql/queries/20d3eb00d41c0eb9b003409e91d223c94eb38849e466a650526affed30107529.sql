  CREATE TABLE test_on_fly_mutations (id UInt64, v String)
  ENGINE = MergeTree ORDER BY id;

  -- تعطيل materialization في الخلفية لعمليات mutation لعرض
  -- السلوك الافتراضي عندما لا تكون on-the-fly mutations مُمكّنة
  SYSTEM STOP MERGES test_on_fly_mutations;
  SET mutations_sync = 0;

  -- إدراج بعض الصفوف في جدولنا الجديد
  INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

  -- تحديث قيم الصفوف
  ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
  ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
  ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';
