-- 実行可能UDFを使用する
SELECT
    review_text,
    sentiment_score(review_text) AS score
FROM customer_reviews;
