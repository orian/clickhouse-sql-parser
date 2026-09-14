-- 매일 오전 2시에 자동으로 갱신됩니다
CREATE MATERIALIZED VIEW monthly_sales_report
REFRESH EVERY 1 DAY OFFSET 2 HOUR
AS SELECT
    toStartOfMonth(order_date) AS month,
    region,
    product_category,
    count() AS order_count,
    sum(amount) AS total_revenue,
    avg(amount) AS avg_order_value
FROM orders
WHERE order_date >= today() - INTERVAL 13 MONTH
GROUP BY month, region, product_category;

-- 쿼리는 항상 최신 데이터를 반환합니다
SELECT * FROM monthly_sales_report
WHERE month = toStartOfMonth(today());
