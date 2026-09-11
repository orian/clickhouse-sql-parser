  SELECT
      sum(total_votes),
      product_category
  FROM amazon.amazon_reviews
  GROUP BY product_category
  ORDER BY 1 DESC
