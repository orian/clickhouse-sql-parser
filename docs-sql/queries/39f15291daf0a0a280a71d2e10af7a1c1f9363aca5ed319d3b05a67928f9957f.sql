SET allow_experimental_ai_functions = 1;
CREATE TABLE product_reviews (review String) ENGINE = Memory;
INSERT INTO product_reviews VALUES ('It works well under rain.');
SELECT review FROM product_reviews ORDER BY aiSimilarity(review, 'It works well under rain', 'text-embedding-3-small') DESC LIMIT 100
