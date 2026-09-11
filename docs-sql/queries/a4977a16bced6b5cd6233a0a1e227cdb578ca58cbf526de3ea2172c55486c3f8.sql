-- 对于不带元组的单个参数，函数将直接返回该参数本身作为 Hilbert 索引，因为无需进行维度映射。
SELECT hilbertEncode(1)
