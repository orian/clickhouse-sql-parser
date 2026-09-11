-- Definir em CREATE TABLE
CREATE TABLE events
(
  ts DateTime,
  user_id UInt64,
  value UInt32,
  INDEX ts_minmax ts TYPE minmax GRANULARITY 1
)
ENGINE=MergeTree
ORDER BY ts;

-- Ou adicionar depois e materializar
ALTER TABLE events ADD INDEX ts_minmax ts TYPE minmax GRANULARITY 1;
ALTER TABLE events MATERIALIZE INDEX ts_minmax;

-- Consulta que se beneficia do índice
SELECT count() FROM events WHERE ts >= now() - 3600;

-- Verificar uso
EXPLAIN indexes = 1
SELECT count() FROM events WHERE ts >= now() - 3600;
