
        INSERT INTO t_low_card_str
        SELECT
            concat('k8_____', toString(number % 4)),
            concat('k16____________', toString(number % 4)),
            concat('k24____________________', toString(number % 4)),
            concat('k32____________________________', toString(number % 4))
        FROM numbers_mt(50000000)
    