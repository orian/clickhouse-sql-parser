SET allow_experimental_ai_functions = 1;
CREATE TABLE reviews (review String) ENGINE = Memory;
INSERT INTO reviews VALUES ('The screen is bright, but the battery lasts only two hours.');
SELECT aiExtract(review, '{"sentiment": "positive, negative or neutral", "topic": "main topic of the review"}') FROM reviews LIMIT 5
