
      CREATE TABLE tab (
          id UInt64,
          str String,
      ) ENGINE = MergeTree()
      ORDER BY id
  