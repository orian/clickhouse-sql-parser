SET allow_experimental_ai_functions = 1;
CREATE TABLE articles (title String) ENGINE = Memory;
INSERT INTO articles VALUES ('ClickHouse is a fast analytical database.');
SELECT aiEmbed(title, 'text-embedding-3-small', map('credentials', 'ai_embedding_credentials', 'dimensions', '256')) FROM articles LIMIT 10
