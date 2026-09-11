EXPLAIN indexes = 1
SELECT
   max(price)
FROM
   uk.uk_price_paid_simple
WHERE town = 'LONDON';
