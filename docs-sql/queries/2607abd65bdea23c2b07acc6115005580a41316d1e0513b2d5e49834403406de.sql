CREATE TABLE encryption_test
(
    `comment` String,
    `secret` String
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO encryption_test VALUES
('aes-256-ofb no IV', encrypt('aes-256-ofb', 'Secret', '12345678910121314151617181920212')),
('aes-256-ofb no IV, different key', encrypt('aes-256-ofb', 'Secret', 'keykeykeykeykeykeykeykeykeykeyke')),
('aes-256-ofb with IV', encrypt('aes-256-ofb', 'Secret', '12345678910121314151617181920212', 'iviviviviviviviv')),
('aes-256-cbc no IV', encrypt('aes-256-cbc', 'Secret', '12345678910121314151617181920212'));

SELECT comment, hex(secret) FROM encryption_test;
