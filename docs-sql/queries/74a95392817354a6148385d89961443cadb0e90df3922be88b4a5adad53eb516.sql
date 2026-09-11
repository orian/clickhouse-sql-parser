  CREATE IF NOT EXISTS test_db;
         
  -- إنشاء جدول مُقسّم
  CREATE TABLE test_db.partitioned (
      id UInt32,
      data String,
      partition_key UInt8
  ) ENGINE = MergeTree()
  PARTITION BY partition_key
  ORDER BY id;

  INSERT INTO test_db.partitioned VALUES
  (1, 'data1', 1),
  (2, 'data2', 2),
  (3, 'data3', 3),
  (4, 'data4', 4);

  SELECT count() FROM test_db.partitioned;

  SELECT partition_key, count() 
  FROM test_db.partitioned
  GROUP BY partition_key
  ORDER BY partition_key;
