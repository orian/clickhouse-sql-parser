SELECT crypto_name FROM holdings
EXCEPT DISTINCT
SELECT crypto_name FROM crypto_prices
WHERE price < 10;
