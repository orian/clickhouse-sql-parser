-- Origin SQL:
CREATE INDEX idx ON tab (col0 DESC);
CREATE INDEX idx ON tab col0;
CREATE INDEX IF NOT EXISTS i ON db.t (a) TYPE bloom_filter(0.01) GRANULARITY 3;
CREATE INDEX i ON t a + b TYPE minmax;
CREATE INDEX i ON t ON CLUSTER c (a DESC, b ASC) TYPE minmax;
CREATE UNIQUE INDEX i ON t (a);
CREATE INDEX i ON t (s) TYPE text(tokenizer = 'splitByNonAlpha');
DROP INDEX i ON t;
DROP INDEX IF EXISTS i ON db.t ON CLUSTER c;


-- Format SQL:
CREATE INDEX idx ON tab (col0 DESC);
CREATE INDEX idx ON tab col0;
CREATE INDEX IF NOT EXISTS i ON db.t (a) TYPE bloom_filter(0.01) GRANULARITY 3;
CREATE INDEX i ON t a + b TYPE minmax;
CREATE INDEX i ON t ON CLUSTER c (a DESC, b ASC) TYPE minmax;
CREATE UNIQUE INDEX i ON t (a);
CREATE INDEX i ON t (s) TYPE text(tokenizer = 'splitByNonAlpha');
DROP INDEX i ON t;
DROP INDEX IF EXISTS i ON db.t ON CLUSTER c;
