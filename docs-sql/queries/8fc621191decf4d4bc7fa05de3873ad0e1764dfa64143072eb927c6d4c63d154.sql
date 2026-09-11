  DROP TABLE IF EXISTS table_to_drop
  SYNC SETTINGS max_table_size_to_drop=2000000000000 -- 한도를 2TB로 늘립니다
