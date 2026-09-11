SELECT
    customer_id,
    count() AS total,
    any(review_headline)
FROM amazon_reviews
GROUP BY customer_id
ORDER BY total DESC
FORMAT Null
