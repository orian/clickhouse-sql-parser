-- user_id は一意のユーザー ID、encrypted は暗号化された文字列フィールド、iv は復号/暗号化のための初期ベクトルです。
-- ユーザーは自分の id と暗号化フィールドを復号するための鍵を知っていると仮定します:
CREATE TABLE decrypt_null
(
    dt DateTime,
    user_id UInt32,
    encrypted String,
    iv String
)
ENGINE = MergeTree;

-- データを挿入する:
INSERT INTO decrypt_null VALUES
('2022-08-02 00:00:00', 1, encrypt('aes-256-gcm', 'value1', 'keykeykeykeykeykeykeykeykeykey01', 'iv1'), 'iv1'),
('2022-09-02 00:00:00', 2, encrypt('aes-256-gcm', 'value2', 'keykeykeykeykeykeykeykeykeykey02', 'iv2'), 'iv2'),
('2022-09-02 00:00:01', 3, encrypt('aes-256-gcm', 'value3', 'keykeykeykeykeykeykeykeykeykey03', 'iv3'), 'iv3');

-- 1 つの鍵で復号を試みる
SELECT
    dt,
    user_id,
    tryDecrypt('aes-256-gcm', encrypted, 'keykeykeykeykeykeykeykeykeykey02', iv) AS value
FROM decrypt_null
ORDER BY user_id ASC
