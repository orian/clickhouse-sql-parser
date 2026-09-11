
        INSERT INTO nested_events
        SELECT tuple(
            toInt64(1783987200000 + (number % 86000) * 1000),
            toInt64(cityHash64(number) % 1000),
            true,
            map('geo', [toNullable(if(cityHash64(number, 1) % 4 = 0, 'DE', 'US'))], 'ipGeo', [toNullable('US')]),
            arrayMap(si -> tuple(
                toNullable(concat('bid_', toString(cityHash64(number, si) % 1000))),
                arrayMap(di -> tuple(
                    toNullable(true),
                    arrayMap(bi -> toNullable(toInt64(cityHash64(number, si, di, bi) % 50)), range(intDiv(cityHash64(number, si, di, 11) % 4, 3))),
                    toNullable('USD'),
                    toNullable(toInt64(cityHash64(number, si, di, 1) % 100)),
                    toNullable(if(cityHash64(number, si, di, 2) % 50000 = 0, 'target_deal', concat('deal_', toString(cityHash64(number, si, di, 2) % 300)))),
                    arrayMap(wi -> toNullable(concat('seat_', toString(cityHash64(number, si, di, wi, 6) % 20))), range(intDiv(cityHash64(number, si, di, 6) % 4, 3))),
                    toNullable(concat('terms_', toString(cityHash64(number, si, di, 7) % 12), '_', repeat('x', 200)))
                ), range(cityHash64(number, si, 8) % 3)),
                toNullable(toInt64(cityHash64(number, si, 9) % 10000))
            ), range(1 + number % 2))
        )
        FROM numbers_mt(4000000)
        SETTINGS min_insert_block_size_rows = 10000, max_insert_block_size = 10000, max_insert_threads = 4
    