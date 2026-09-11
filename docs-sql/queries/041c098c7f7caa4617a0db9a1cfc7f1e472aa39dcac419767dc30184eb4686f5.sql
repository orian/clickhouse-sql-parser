-- 较低优先级：对所有敏感数据进行基础脱敏
CREATE MASKING POLICY basic_masking ON orders
    UPDATE
        name = '****',
        email = '****@****.com'
    TO masked_data_viewer
    PRIORITY 0;

-- 较高优先级：更精细的脱敏处理（覆盖 basic_masking）
CREATE MASKING POLICY refined_masking ON orders
    UPDATE
        name = replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****')
    WHERE total_amount > 100
    TO masked_data_viewer
    PRIORITY 10;
