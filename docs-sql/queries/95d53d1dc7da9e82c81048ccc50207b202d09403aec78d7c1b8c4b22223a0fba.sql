  SELECT
      any(product_title),
      count()
  FROM amazon.amazon_reviews
  GROUP BY product_id
  ORDER BY 2 DESC
  LIMIT 10;
