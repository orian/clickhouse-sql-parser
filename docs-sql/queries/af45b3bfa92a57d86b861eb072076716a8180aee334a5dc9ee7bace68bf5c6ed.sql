CREATE MASKING POLICY hash_sensitive_data ON orders
    UPDATE
        email = concat(toString(cityHash64(email)), '@masked.com'),
        phone = concat('555-', toString(cityHash64(phone) % 10000000))
    TO masked_data_viewer;
