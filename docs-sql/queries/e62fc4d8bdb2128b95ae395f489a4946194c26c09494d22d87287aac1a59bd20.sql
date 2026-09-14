  SELECT MAX(price) AS highest_price
  FROM uk.uk_price_paid_simple_partitioned
  WHERE date >= '2020-12-01'
    AND date <= '2020-12-31'
    AND town = 'LONDON';
