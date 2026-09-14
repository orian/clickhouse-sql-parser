   -- 按年份显示所有 uk_price_paid 交易的总价格和总笔数。
   SELECT year(date), sum(price) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
