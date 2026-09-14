CREATE TABLE user_data (
user_id UInt32,
first_name String,
last_name String,
score UInt32
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO user_data VALUES
(1, 'John', 'Doe', 100),
(2, 'Jane', 'Smith', 150),
(3, 'John', 'Wilson', 120),
(4, 'Jane', 'Smith', 90);

SELECT
    concatWithSeparatorAssumeInjective('-', first_name, last_name) as full_name,
    sum(score) as total_score
FROM user_data
GROUP BY concatWithSeparatorAssumeInjective('-', first_name, last_name);
