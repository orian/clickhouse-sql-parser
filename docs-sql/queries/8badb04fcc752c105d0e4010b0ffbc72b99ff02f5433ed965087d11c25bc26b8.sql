  INSERT INTO hierarchy_source VALUES
  (0, 0, 'Root'),
  (1, 0, 'Level 1 - Node 1'),
  (2, 1, 'Level 2 - Node 2'),
  (3, 1, 'Level 2 - Node 3'),
  (4, 2, 'Level 3 - Node 4'),
  (5, 2, 'Level 3 - Node 5'),
  (6, 3, 'Level 3 - Node 6');

  -- 0 (Raíz)
  -- └── 1 (Nivel 1 - Nodo 1)
  --     ├── 2 (Nivel 2 - Nodo 2)
  --     │   ├── 4 (Nivel 3 - Nodo 4)
  --     │   └── 5 (Nivel 3 - Nodo 5)
  --     └── 3 (Nivel 2 - Nodo 3)
  --         └── 6 (Nivel 3 - Nodo 6)
