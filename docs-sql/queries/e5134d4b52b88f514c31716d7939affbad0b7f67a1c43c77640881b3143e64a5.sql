CREATE VIEW masked_orders AS
SELECT
    user_id,
    replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****') AS name,
    replaceRegexpOne(email, '^(.{0})[^@]*(@.*)$', '\\1****\\2') AS email,
    replaceRegexpOne(phone, '^(\\d{3})-(\\d{3})-(\\d{4})$', '\\1-***-\\3') AS phone,
    total_amount,
    order_date,
    replaceRegexpOne(shipping_address, '^[^,]+,\\s*(.*)$', '*** \\1') AS shipping_address
FROM orders;
