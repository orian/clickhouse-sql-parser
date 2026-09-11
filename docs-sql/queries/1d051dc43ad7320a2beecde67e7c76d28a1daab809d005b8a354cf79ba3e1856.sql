   -- Montrez-moi le prix total et le nombre total de toutes les transactions uk_price_paid par année.
   SELECT year(date), sum(price) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
