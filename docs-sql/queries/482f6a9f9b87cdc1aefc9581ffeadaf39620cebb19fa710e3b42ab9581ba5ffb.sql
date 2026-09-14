   -- Affiche le prix total et le nombre total de toutes les transactions uk_price_paid pour chaque année.
   SELECT year(date), sum(pricee) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
