EXPLAIN actions = 1
SELECT
    helpful_votes,
    product_title,
    review_headline,
    review_body
FROM amazon.amazon_reviews
ORDER BY helpful_votes DESC
LIMIT 3
SETTINGS
    query_plan_optimize_lazy_materialization = true;
