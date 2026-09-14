-- 创建表并插入数据
CREATE TABLE test_lazy (json JSON) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO test_lazy VALUES ('{"user_id": "123", "score": "95.5"}');

-- 启用实验性设置
SET allow_experimental_json_lazy_type_hints = 1;

-- 添加类型提示 - 此操作无需变更，立即完成
ALTER TABLE test_lazy MODIFY COLUMN json JSON(user_id UInt64, score Float64);

-- 查询数据 - 类型提示在读取时生效
SELECT json.user_id, toTypeName(json.user_id), json.score, toTypeName(json.score) FROM test_lazy;
