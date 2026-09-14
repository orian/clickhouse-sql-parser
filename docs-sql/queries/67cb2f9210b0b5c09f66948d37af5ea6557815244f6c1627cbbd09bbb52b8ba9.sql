  SELECT
      toStartOfMonth(review_date) AS month,
      any(product_title),
      avg(star_rating) AS avg_stars
  FROM amazon.amazon_reviews
  GROUP BY
      month,
      product_id
  ORDER BY
      month DESC,
      product_id ASC
  LIMIT 20;
