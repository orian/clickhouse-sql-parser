-- 샘플 데이터: 사용자가 결제 시작 -> 상품 상세 -> 결제 재시작 -> 결제 완료 순으로 방문
-- allow_reentry 미사용 시: 2단계(상품 상세 페이지)에서 중단
-- allow_reentry 사용 시: 4단계(결제 완료)까지 도달

SELECT
    level,
    count() AS users
FROM
(
    SELECT
        user_id,
        windowFunnel(3600, 'strict_order', 'allow_reentry')(
            timestamp,
            action = 'begin_checkout',      -- 1단계: 결제 시작
            action = 'view_product_detail', -- 2단계: 상품 상세 조회
            action = 'begin_checkout',      -- 3단계: 결제 재시작 (재진입)
            action = 'complete_payment'     -- 4단계: 결제 완료
        ) AS level
    FROM user_events
    WHERE event_date = today()
    GROUP BY user_id
)
GROUP BY level
ORDER BY level ASC;
