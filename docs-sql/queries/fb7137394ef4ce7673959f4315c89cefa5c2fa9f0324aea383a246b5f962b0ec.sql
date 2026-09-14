   -- 年ごとの uk_price_paid の全取引について、合計価格と総件数を表示します。
   SELECT year(date), sum(pricee) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
