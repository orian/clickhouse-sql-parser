SET allow_experimental_ai_functions = 1;
CREATE TABLE articles (body String) ENGINE = Memory;
INSERT INTO articles VALUES ('ClickHouse processes analytical queries quickly.');
SELECT aiTranslate(body, 'Japanese', map('instructions', 'Use polite form (desu/masu)')) FROM articles LIMIT 5
