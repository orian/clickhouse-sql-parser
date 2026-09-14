  INSERT INTO hierarchy_source VALUES
  (0, 0, 'Root'),
  (1, 0, 'Level 1 - Node 1'),
  (2, 1, 'Level 2 - Node 2'),
  (3, 1, 'Level 2 - Node 3'),
  (4, 2, 'Level 3 - Node 4'),
  (5, 2, 'Level 3 - Node 5'),
  (6, 3, 'Level 3 - Node 6');

  -- 0 (Raiz)
  -- └── 1 (Nível 1 - Nó 1)
  --     ├── 2 (Nível 2 - Nó 2)
  --     │   ├── 4 (Nível 3 - Nó 4)
  --     │   └── 5 (Nível 3 - Nó 5)
  --     └── 3 (Nível 2 - Nó 3)
  --         └── 6 (Nível 3 - Nó 6)
