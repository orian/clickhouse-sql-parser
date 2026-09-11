  INSERT INTO hierarchy_source VALUES
  (0, 0, 'Root'),
  (1, 0, 'Level 1 - Node 1'),
  (2, 1, 'Level 2 - Node 2'),
  (3, 1, 'Level 2 - Node 3'),
  (4, 2, 'Level 3 - Node 4'),
  (5, 2, 'Level 3 - Node 5'),
  (6, 3, 'Level 3 - Node 6');

  -- 0 (Racine)
  -- └── 1 (Niveau 1 - Nœud 1)
  --     ├── 2 (Niveau 2 - Nœud 2)
  --     │   ├── 4 (Niveau 3 - Nœud 4)
  --     │   └── 5 (Niveau 3 - Nœud 5)
  --     └── 3 (Niveau 2 - Nœud 3)
  --         └── 6 (Niveau 3 - Nœud 6)
