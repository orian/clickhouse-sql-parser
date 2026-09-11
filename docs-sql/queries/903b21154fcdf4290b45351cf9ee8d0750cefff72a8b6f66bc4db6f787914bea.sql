ALTER TABLE t MATERIALIZE INDEX idx_bf;

EXPLAIN indexes = 1
SELECT count() FROM t WHERE u64 IN (123, 456);

-- Opcional: informações detalhadas de pruning
SET send_logs_level = 'trace';
