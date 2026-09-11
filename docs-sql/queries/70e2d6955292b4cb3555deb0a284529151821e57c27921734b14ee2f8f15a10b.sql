-- 낮은 우선순위: 모든 민감한 데이터에 대한 기본 마스킹
CREATE MASKING POLICY basic_masking ON orders
    UPDATE
        name = '****',
        email = '****@****.com'
    TO masked_data_viewer
    PRIORITY 0;

-- 높은 우선순위: 더 정교한 마스킹 (basic_masking 재정의)
CREATE MASKING POLICY refined_masking ON orders
    UPDATE
        name = replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****')
    WHERE total_amount > 100
    TO masked_data_viewer
    PRIORITY 10;
