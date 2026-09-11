-- タプルなしの単一引数の場合、次元マッピングが不要なため、関数は引数自体を
-- Hilbert 索引としてそのまま返します。
SELECT hilbertEncode(1)
