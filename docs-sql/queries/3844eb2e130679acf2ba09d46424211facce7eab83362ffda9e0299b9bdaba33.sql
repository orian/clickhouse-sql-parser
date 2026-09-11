SELECT crypto_name FROM holdings
INTERSECT
SELECT crypto_name FROM crypto_prices
WHERE price > 100
