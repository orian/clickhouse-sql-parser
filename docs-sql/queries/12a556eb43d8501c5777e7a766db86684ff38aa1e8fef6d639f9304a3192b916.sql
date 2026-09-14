  INSERT INTO hierarchy_source VALUES
  (0, 0, 'Root'),
  (1, 0, 'Level 1 - Node 1'),
  (2, 1, 'Level 2 - Node 2'),
  (3, 1, 'Level 2 - Node 3'),
  (4, 2, 'Level 3 - Node 4'),
  (5, 2, 'Level 3 - Node 5'),
  (6, 3, 'Level 3 - Node 6');

  -- 0 (루트)
  -- └── 1 (수준 1 - 노드 1)
  --     ├── 2 (수준 2 - 노드 2)
  --     │   ├── 4 (수준 3 - 노드 4)
  --     │   └── 5 (수준 3 - 노드 5)
  --     └── 3 (수준 2 - 노드 3)
  --         └── 6 (수준 3 - 노드 6)
