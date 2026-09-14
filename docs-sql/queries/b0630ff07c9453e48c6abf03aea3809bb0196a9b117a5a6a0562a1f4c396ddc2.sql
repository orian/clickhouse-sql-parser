SELECT * FROM
(
    SELECT customer, sum(amount) AS total FROM
    (
        SELECT * FROM
        (
            SELECT * FROM orders
        )
        WHERE cancelled = 0
    )
    GROUP BY customer
)
ORDER BY total DESC
LIMIT 3
