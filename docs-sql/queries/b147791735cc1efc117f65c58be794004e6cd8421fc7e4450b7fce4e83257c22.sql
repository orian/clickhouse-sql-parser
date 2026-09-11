-- user_id는 고유한 사용자 ID, encrypted는 암호화된 문자열 필드, iv는 복호화/암호화를 위한 초기 벡터입니다.
-- 사용자가 자신의 ID와 encrypted 필드를 복호화하기 위한 키를 알고 있다고 가정합니다:
CREATE TABLE decrypt_null
(
    dt DateTime,
    user_id UInt32,
    encrypted String,
    iv String
)
ENGINE = MergeTree;

-- 데이터를 삽입합니다:
INSERT INTO decrypt_null VALUES
('2022-08-02 00:00:00', 1, encrypt('aes-256-gcm', 'value1', 'keykeykeykeykeykeykeykeykeykey01', 'iv1'), 'iv1'),
('2022-09-02 00:00:00', 2, encrypt('aes-256-gcm', 'value2', 'keykeykeykeykeykeykeykeykeykey02', 'iv2'), 'iv2'),
('2022-09-02 00:00:01', 3, encrypt('aes-256-gcm', 'value3', 'keykeykeykeykeykeykeykeykeykey03', 'iv3'), 'iv3');

-- 하나의 키로 복호화를 시도합니다
SELECT
    dt,
    user_id,
    tryDecrypt('aes-256-gcm', encrypted, 'keykeykeykeykeykeykeykeykeykey02', iv) AS value
FROM decrypt_null
ORDER BY user_id ASC
