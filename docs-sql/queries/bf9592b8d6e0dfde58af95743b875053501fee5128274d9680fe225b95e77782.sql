  DROP TABLE IF EXISTS table_to_drop
  SYNC SETTINGS max_table_size_to_drop=2000000000000 -- 将限制提高到 2TB
