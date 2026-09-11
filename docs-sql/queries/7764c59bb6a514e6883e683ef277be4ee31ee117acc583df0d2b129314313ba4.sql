-- 创建一个表，其中 user_id 是唯一用户 ID，encrypted 是加密字符串字段，iv 是用于解密/加密的初始向量。
-- 假设用户知道自己的 ID 以及用于解密 encrypted 字段的密钥：
CREATE TABLE decrypt_null
(
    dt DateTime,
    user_id UInt32,
    encrypted String,
    iv String
)
ENGINE = MergeTree;

-- 插入一些数据：
INSERT INTO decrypt_null VALUES
('2022-08-02 00:00:00', 1, encrypt('aes-256-gcm', 'value1', 'keykeykeykeykeykeykeykeykeykey01', 'iv1'), 'iv1'),
('2022-09-02 00:00:00', 2, encrypt('aes-256-gcm', 'value2', 'keykeykeykeykeykeykeykeykeykey02', 'iv2'), 'iv2'),
('2022-09-02 00:00:01', 3, encrypt('aes-256-gcm', 'value3', 'keykeykeykeykeykeykeykeykeykey03', 'iv3'), 'iv3');

-- 尝试使用一个密钥解密
SELECT
    dt,
    user_id,
    tryDecrypt('aes-256-gcm', encrypted, 'keykeykeykeykeykeykeykeykeykey02', iv) AS value
FROM decrypt_null
ORDER BY user_id ASC
