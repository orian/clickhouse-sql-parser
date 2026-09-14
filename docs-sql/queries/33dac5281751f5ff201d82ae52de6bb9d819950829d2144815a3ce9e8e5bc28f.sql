    DELIMITER $$

    CREATE PROCEDURE process_order(
        IN p_order_id INT,
        IN p_customer_id INT,
        IN p_order_total DECIMAL(10,2),
        OUT p_status VARCHAR(50),
        OUT p_loyalty_points INT
    )
    BEGIN
        DECLARE v_customer_tier VARCHAR(20);
        DECLARE v_previous_orders INT;
        DECLARE v_discount DECIMAL(10,2);

        -- 트랜잭션 시작
        START TRANSACTION;

        -- 고객 정보 조회
        SELECT tier, total_orders
        INTO v_customer_tier, v_previous_orders
        FROM customers
        WHERE customer_id = p_customer_id;

        -- 등급에 따른 할인 계산
        IF v_customer_tier = 'gold' THEN
            SET v_discount = p_order_total * 0.15;
        ELSEIF v_customer_tier = 'silver' THEN
            SET v_discount = p_order_total * 0.10;
        ELSE
            SET v_discount = 0;
        END IF;

        -- 주문 레코드 삽입
        INSERT INTO orders (order_id, customer_id, order_total, discount, final_amount)
        VALUES (p_order_id, p_customer_id, p_order_total, v_discount,
                p_order_total - v_discount);

        -- 고객 통계 업데이트
        UPDATE customers
        SET total_orders = total_orders + 1,
            lifetime_value = lifetime_value + (p_order_total - v_discount),
            last_order_date = NOW()
        WHERE customer_id = p_customer_id;

        -- 적립 포인트 계산 (1달러당 1포인트)
        SET p_loyalty_points = FLOOR(p_order_total - v_discount);

        -- 적립 포인트 트랜잭션 삽입
        INSERT INTO loyalty_points (customer_id, points, transaction_date, description)
        VALUES (p_customer_id, p_loyalty_points, NOW(),
                CONCAT('Order #', p_order_id));

        -- 고객 등급 업그레이드 여부 확인
        IF v_previous_orders + 1 >= 10 AND v_customer_tier = 'bronze' THEN
            UPDATE customers SET tier = 'silver' WHERE customer_id = p_customer_id;
            SET p_status = 'ORDER_COMPLETE_TIER_UPGRADED_SILVER';
        ELSEIF v_previous_orders + 1 >= 50 AND v_customer_tier = 'silver' THEN
            UPDATE customers SET tier = 'gold' WHERE customer_id = p_customer_id;
            SET p_status = 'ORDER_COMPLETE_TIER_UPGRADED_GOLD';
        ELSE
            SET p_status = 'ORDER_COMPLETE';
        END IF;

        COMMIT;
    END$$

    DELIMITER ;

    -- 저장 프로시저 호출
    CALL process_order(12345, 5678, 250.00, @status, @points);
    SELECT @status, @points;
