   -- اعرض لي إجمالي السعر وإجمالي عدد جميع معاملات uk_price_paid حسب السنة.
   SELECT year(date), sum(pricee) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
