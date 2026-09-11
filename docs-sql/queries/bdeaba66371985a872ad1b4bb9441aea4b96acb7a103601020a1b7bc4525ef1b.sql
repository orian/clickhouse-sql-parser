SELECT crypto_name FROM holdings
INTERSECT DISTINCT
SELECT crypto_name FROM crypto_prices
WHERE price > 100;
