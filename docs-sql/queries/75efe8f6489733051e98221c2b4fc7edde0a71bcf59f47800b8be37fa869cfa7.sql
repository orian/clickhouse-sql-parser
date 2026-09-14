-- 当参数的取值范围（或基数）差异悬殊时，范围扩展有助于获得相近的分布
-- 例如：'IP Address'（0...FFFFFFFF）与 'Country code'（0...FF）
-- 注意：Tuple 的大小必须等于其他参数的数量。
SELECT mortonEncode((1,2), 1024, 16)
