  INSERT INTO hierarchy_source VALUES
  (0, 0, 'Root'),
  (1, 0, 'Level 1 - Node 1'),
  (2, 1, 'Level 2 - Node 2'),
  (3, 1, 'Level 2 - Node 3'),
  (4, 2, 'Level 3 - Node 4'),
  (5, 2, 'Level 3 - Node 5'),
  (6, 3, 'Level 3 - Node 6');

  -- 0（根）
  -- └── 1（第 1 层 - 节点 1）
  --     ├── 2（第 2 层 - 节点 2）
  --     │   ├── 4（第 3 层 - 节点 4）
  --     │   └── 5（第 3 层 - 节点 5）
  --     └── 3（第 2 层 - 节点 3）
  --         └── 6（第 3 层 - 节点 6）
