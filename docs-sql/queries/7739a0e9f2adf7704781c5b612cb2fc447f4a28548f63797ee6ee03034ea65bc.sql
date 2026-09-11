  CREATE DATABASE test_db;

  CREATE TABLE test_db.test_table (
      id UUID,
      name String,
      email String,
      age UInt8,
      salary UInt32,
      created_at DateTime,
      is_active UInt8,
      department String,
      score Float32,
      country String
  ) ENGINE = MergeTree()
  ORDER BY id;
