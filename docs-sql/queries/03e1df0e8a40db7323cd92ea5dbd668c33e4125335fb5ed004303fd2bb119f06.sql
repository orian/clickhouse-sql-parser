-- 当参数的取值范围（或基数）差异悬殊时，范围扩展有助于实现相近的分布。
-- 例如：'IP Address'（0...FFFFFFFF）与 'Country code'（0...FF）。
-- 注意：Tuple 的大小必须与其他参数的数量相等。
SELECT hilbertEncode((10, 6), 1024, 16)
