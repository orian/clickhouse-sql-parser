   -- Mostre o preço total e o número total de todas as transações uk_price_paid por ano.
   SELECT year(date), sum(pricee) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
