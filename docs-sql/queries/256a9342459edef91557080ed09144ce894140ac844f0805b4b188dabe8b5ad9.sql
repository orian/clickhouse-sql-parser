SELECT MAX(price) AS highest_price
FROM uk.uk_price_paid_simple_partitioned
WHERE town = 'LONDON';
