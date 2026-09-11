ALTER TABLE t MATERIALIZE INDEX idx_bf;

EXPLAIN indexes = 1
SELECT count() FROM t WHERE u64 IN (123, 456);

-- 선택 사항: 상세 프루닝 정보
SET send_logs_level = 'trace';
