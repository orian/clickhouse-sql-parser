CREATE FUNCTION bfEstimateFunctions AS
(total_grams, bits) -> round((bits / total_grams) * log(2));

CREATE FUNCTION bfEstimateBmSize AS
(total_grams, p_false) -> ceil((total_grams * log(p_false)) / log(1 / pow(2, log(2))));

-- 4300 N-gramのサイズ計算例、p_false = 0.0001
SELECT bfEstimateBmSize(4300, 0.0001) / 8 AS size_bytes;  -- 約10304
SELECT bfEstimateFunctions(4300, bfEstimateBmSize(4300, 0.0001)) AS k; -- 約13
