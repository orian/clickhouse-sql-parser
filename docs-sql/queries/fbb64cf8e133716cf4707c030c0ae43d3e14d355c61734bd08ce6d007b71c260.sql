CREATE TABLE articles (article_title String, article_body String) ENGINE = Memory;
INSERT INTO articles VALUES ('ClickHouse', 'ClickHouse is an open-source column-oriented database for online analytical processing.');
SELECT article_title, aiGenerate(concat('Summarize in one sentence: ', article_body)) AS summary FROM articles LIMIT 5
