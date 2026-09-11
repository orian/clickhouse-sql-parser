CREATE FUNCTION bfEstimateFunctions AS
(total_grams, bits) -> round((bits / total_grams) * log(2));

CREATE FUNCTION bfEstimateBmSize AS
(total_grams, p_false) -> ceil((total_grams * log(p_false)) / log(1 / pow(2, log(2))));

-- 4300개의 ngram, p_false = 0.0001에 대한 크기 계산 예시
SELECT bfEstimateBmSize(4300, 0.0001) / 8 AS size_bytes;  -- ~10304 (바이트 단위 크기)
SELECT bfEstimateFunctions(4300, bfEstimateBmSize(4300, 0.0001)) AS k; -- ~13 (해시 함수 수)
