-- 低優先度: すべての機密データに対する基本マスキング
CREATE MASKING POLICY basic_masking ON orders
    UPDATE
        name = '****',
        email = '****@****.com'
    TO masked_data_viewer
    PRIORITY 0;

-- 高優先度: より精緻なマスキング（basic_maskingを上書き）
CREATE MASKING POLICY refined_masking ON orders
    UPDATE
        name = replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****')
    WHERE total_amount > 100
    TO masked_data_viewer
    PRIORITY 10;
