   -- 显示 uk_price_paid 中按年份统计的交易总价和交易总数。
   SELECT year(date), sum(pricee) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
