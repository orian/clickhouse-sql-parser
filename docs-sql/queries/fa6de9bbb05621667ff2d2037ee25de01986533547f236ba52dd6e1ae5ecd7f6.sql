  SELECT toYear(date) AS year, avg(price) AS avg_price
  FROM uk_price_paid
  GROUP BY year
  ORDER BY year;
