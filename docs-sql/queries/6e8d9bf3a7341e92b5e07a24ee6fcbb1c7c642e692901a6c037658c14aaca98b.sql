SET param_ids = [1, 2, 3, 4, 5];

SELECT * FROM products WHERE id IN {ids: Array(UInt32)};
