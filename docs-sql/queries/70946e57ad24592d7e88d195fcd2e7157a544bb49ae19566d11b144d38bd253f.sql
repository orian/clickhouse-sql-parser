INSERT INTO default.orders
SELECT * FROM lake.`sales.orders`;
