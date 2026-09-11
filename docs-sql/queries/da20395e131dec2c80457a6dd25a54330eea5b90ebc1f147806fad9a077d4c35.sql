CREATE MASKING POLICY mask_high_value_orders ON orders
    UPDATE
        name = replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****'),
        email = replaceRegexpOne(email, '^(.{2})[^@]*(@.*)$', '\\1****\\2')
    WHERE total_amount > 200
    TO masked_data_viewer;
