CREATE TABLE hierarchy_source (id UInt64, parent_id UInt64, name String) ENGINE = Memory;
INSERT INTO hierarchy_source VALUES (0, 0, 'Root'), (1, 0, 'Level 1 - Node 1'), (2, 1, 'Level 2 - Node 2'), (3, 1, 'Level 2 - Node 3'), (4, 2, 'Level 3 - Node 4'), (5, 2, 'Level 3 - Node 5'), (6, 3, 'Level 3 - Node 6');
CREATE DICTIONARY hierarchical_dictionary (id UInt64, parent_id UInt64 HIERARCHICAL, name String) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'hierarchy_source' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- consider the following hierarchical dictionary:
-- 0 (Root)
-- └── 1 (Level 1 - Node 1)
--     ├── 2 (Level 2 - Node 2)
--     │   ├── 4 (Level 3 - Node 4)
--     │   └── 5 (Level 3 - Node 5)
--     └── 3 (Level 2 - Node 3)
--         └── 6 (Level 3 - Node 6)

SELECT dictHas('hierarchical_dictionary', 2);
SELECT dictHas('hierarchical_dictionary', 7);
