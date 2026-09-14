-- Usar el UDF ejecutable
SELECT
    review_text,
    sentiment_score(review_text) AS score
FROM customer_reviews;
