SET allow_experimental_ai_functions = 1;
CREATE TABLE docs (id UInt64, title String) ENGINE = Memory;
INSERT INTO docs VALUES (1, 'ClickHouse documentation'), (2, 'ClickHouse database guide');
SELECT a.id, b.id FROM docs a, docs b WHERE a.id < b.id AND aiSimilarity(a.title, b.title, 'text-embedding-3-small') > 0.9
