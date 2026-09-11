   -- Muéstrame el precio total y el número total de todas las transacciones de uk_price_paid por año.
   SELECT year(date), sum(price) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
