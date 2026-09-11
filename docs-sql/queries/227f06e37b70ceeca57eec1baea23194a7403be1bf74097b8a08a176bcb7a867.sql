CREATE MASKING POLICY mask_pii_data ON orders
    UPDATE
        name = replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****'),
        email = replaceRegexpOne(email, '^(.{2})[^@]*(@.*)$', '\\1****\\2'),
        phone = replaceRegexpOne(phone, '^(\\d{3})-(\\d{3})-(\\d{4})$', '\\1-***-\\3'),
        shipping_address = replaceRegexpOne(shipping_address, '^[^,]+,\\s*(.*)$', '*** \\1')
    TO masked_data_viewer;
