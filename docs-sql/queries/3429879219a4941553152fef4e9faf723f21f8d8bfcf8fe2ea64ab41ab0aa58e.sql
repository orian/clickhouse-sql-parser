CREATE TABLE users (
    id UInt32,
    name String,
    age UInt8,
    city String
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO users VALUES
(1, 'Alice', 25, 'New York'),
(2, 'Bob', 30, 'London'),
(3, 'Charlie', 35, 'Tokyo');

SELECT groupBitXor(cityHash64(*)) FROM users;
