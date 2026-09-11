  INSERT INTO hierarchy_source VALUES
  (0, 0, 'Root'),
  (1, 0, 'Level 1 - Node 1'),
  (2, 1, 'Level 2 - Node 2'),
  (3, 1, 'Level 2 - Node 3'),
  (4, 2, 'Level 3 - Node 4'),
  (5, 2, 'Level 3 - Node 5'),
  (6, 3, 'Level 3 - Node 6');

  -- 0 (ルート)
  -- └── 1 (レベル 1 - ノード 1)
  --     ├── 2 (レベル 2 - ノード 2)
  --     │   ├── 4 (レベル 3 - ノード 4)
  --     │   └── 5 (レベル 3 - ノード 5)
  --     └── 3 (レベル 2 - ノード 3)
  --         └── 6 (レベル 3 - ノード 6)
