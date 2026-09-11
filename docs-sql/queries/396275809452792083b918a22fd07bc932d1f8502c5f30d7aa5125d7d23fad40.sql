  SELECT
      toYear(date) AS year,
      round(avg(price)) AS price,
      bar(price, 0, 1000000, 80)
  FROM uk.uk_price_paid_with_projections_v2
  GROUP BY year
  ORDER BY year ASC

