   -- 연도별로 모든 uk_price_paid 거래의 총금액과 총건수를 보여줍니다.
   SELECT year(date), sum(pricee) as total_price, Count(*) as total_transactions
   FROM uk_price_paid
   GROUP BY year(date)
