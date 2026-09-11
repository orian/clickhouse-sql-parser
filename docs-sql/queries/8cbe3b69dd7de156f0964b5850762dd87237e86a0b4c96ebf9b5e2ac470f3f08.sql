ALTER TABLE t MATERIALIZE INDEX idx_bf;

EXPLAIN indexes = 1
SELECT count() FROM t WHERE u64 IN (123, 456);

-- 可选：查看详细的剪枝信息
SET send_logs_level = 'trace';
