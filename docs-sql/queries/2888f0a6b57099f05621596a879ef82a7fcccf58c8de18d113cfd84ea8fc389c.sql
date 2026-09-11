   -- 年ごとに、uk_price_paid のすべてのトランザクションについて合計価格と総件数を表示してください。
   SELECT year(date), sum(price) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
