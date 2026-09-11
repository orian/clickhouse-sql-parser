-- 실행형 UDF 사용
SELECT
    review_text,
    sentiment_score(review_text) AS score
FROM customer_reviews;
