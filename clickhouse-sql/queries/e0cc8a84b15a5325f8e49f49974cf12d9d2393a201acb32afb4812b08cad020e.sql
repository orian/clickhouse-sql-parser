
    SELECT *
    FROM
    (
        SELECT
            cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(toUInt128('1015029'))) AS idColumnU64,
            *
        FROM
        (
            SELECT
                if(max(nDateTime_02_date_292929292) > '2020-10-31 00:00:00', max(nDateTime_02_date_292929292), NULL) AS o1,
                if(max(other_max_datetime_05) > '2020-10-31 00:00:00', max(other_max_datetime_05), NULL) AS o2,
                if(max(nDateTime_03_date) > '2020-10-31 00:00:00', max(nDateTime_03_date), NULL) AS o3,
                if(max(nDateTime_04_date) > '2020-10-31 00:00:00', max(nDateTime_04_date), NULL) AS o4,
                if(max(nDateTime_02_date) > '2020-10-31 00:00:00', max(nDateTime_02_date), NULL) AS o5,
                if(max(other_max_datetime_01) > '2020-10-31 00:00:00', max(other_max_datetime_01), NULL) AS o6,
                if(max(other_max_datetime_02) > '2020-10-31 00:00:00', max(other_max_datetime_02), NULL) AS o7,
                argMaxMerge(agg_argmax_string_datetime_13) AS o8,
                argMaxMerge(agg_argmax_string_datetime_05) AS o9,
                argMaxMerge(agg_argmax_string_datetime_06) AS o10,
                argMaxMerge(agg_argmax_string_datetime_02) AS o11,
                argMaxMerge(agg_argmax_string_datetime_04) AS o12,
                argMaxMerge(agg_argmax_string_datetime_15) AS o13,
                argMaxMerge(agg_argmax_string_datetime_01) AS o14,
                argMaxMerge(agg_argmax_string_u8_01) AS o15,
                argMaxMerge(agg_argmax_f32_datetime_02) AS o16,
                if(argMaxMerge(agg_argmax_string_datetime_09) != '', argMaxMerge(agg_argmax_string_datetime_09), NULL) AS o17,
                if(argMaxMerge(agg_argmax_date_datetime_01) > '2020-10-31', argMaxMerge(agg_argmax_date_datetime_01), NULL) AS o18,
                if(argMaxMerge(agg_argmax_date_datetime_02) > '2020-10-31', argMaxMerge(agg_argmax_date_datetime_02), NULL) AS o19,
                argMaxMerge(agg_argmax_u8_other_02) AS o20,
                argMaxMerge(agg_argmax_u8_other_03) AS o21,
                argMaxMerge(agg_argmax_u8_other_04) AS o22,
                argMaxMerge(agg_argmax_u8_other_01) AS o23,
                argMaxMerge(agg_argmax_string_datetime_10) AS o24,
                argMaxMerge(agg_argmax_string_datetime_11) AS o25,
                countMerge(aggCount_3) AS o26,
                countMerge(aggCount_4) AS o27
            FROM interpret_table_01 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(toUInt128('1015029'))) = c.idColumnU64
        ) AS s01,
        (
            WITH (
                    SELECT coalesce(if((topKWeightedMerge(2)(agg_topk_01)[1]) != toUInt128(toUInt128('1015029')), topKWeightedMerge(2)(agg_topk_01)[1], topKWeightedMerge(2)(agg_topk_01)[2]), 0)
                    FROM interpret_table_02 AS c
                    PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2')) = c.idColumnU64
                ) AS other_idU128
            SELECT
                if(max(other_max_datetime_05) > '2020-10-31 00:00:00', max(other_max_datetime_05), NULL) AS o28,
                if(max(other_max_datetime_01) > '2020-10-31 00:00:00', max(other_max_datetime_01), NULL) AS o29,
                if(max(nDateTime_02_date) > '2020-10-31 00:00:00', max(nDateTime_02_date), NULL) AS o30,
                other_idU128
            FROM interpret_table_01 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(other_idU128)) = c.idColumnU64
        ) AS s02,
        (
            SELECT
                minIf(minDate, dateColumn > (now() - toIntervalDay(7))) AS o31,
                maxIf(maxDate, dateColumn > (now() - toIntervalDay(7))) AS o32,
                maxIf(maxInt16, dateColumn > (now() - toIntervalDay(28))) AS o33,
                countMergeIf(aggCount_3, dateColumn > (now() - toIntervalHour(24))) AS o34,
                countMergeIf(aggCount_3, dateColumn > (now() - toIntervalDay(14))) AS o35,
                countMergeIf(aggCount_3, dateColumn > (now() - toIntervalDay(28))) AS o36,
                countMergeIf(aggCount_4, dateColumn > (now() - toIntervalHour(24))) AS o37,
                countMergeIf(aggCount_4, dateColumn > (now() - toIntervalDay(7))) AS o38,
                countMergeIf(aggCount_4, dateColumn > (now() - toIntervalDay(28))) AS o27_month,
                countMergeIf(aggCount_2_shown, dateColumn > (now() - toIntervalDay(14))) AS o40
            FROM interpret_table_03 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(toUInt128('1015029'))) = c.idColumnU64
        ) AS s03,
        (
            SELECT
                countMerge(agg_count_03) AS o41,
                countMerge(agg_count_04) AS o42
            FROM interpret_table_04 AS c
            PREWHERE (cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(toUInt128('1015029'))) = c.idColumnU64) AND (ls_01 = 'exit')
        ) AS s04,
        (
            SELECT
                countMerge(aggCount_3) AS o43,
                countMerge(aggCount_4) AS o44,
                countMerge(aggCount) AS o45
            FROM interpret_table_02 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2')) = c.idColumnU64
        ) AS s05,
        (
            SELECT
                countMergeIf(aggCount_3, dateColumn > (now() - toIntervalDay(14))) AS o46,
                uniqMergeIf(agg_uniq_u128_03, dateColumn > (now() - toIntervalHour(24))) AS o47,
                uniqMergeIf(agg_uniq_u128_03, dateColumn > (now() - toIntervalDay(14))) AS o48,
                countMergeIf(aggCount_4, dateColumn > (now() - toIntervalDay(14))) AS o49,
                countMergeIf(aggCount_4, dateColumn > (now() - toIntervalDay(28))) AS o50
            FROM interpret_table_05 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2')) = c.idColumnU64
        ) AS s06,
        (
            SELECT countMerge(aggCount_3) AS o51
            FROM interpret_table_06 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(joinGet(interpret_table_07, 'idColumnI64', toUInt128('1015029')))) = c.idColumnU64
        ) AS s07,
        (
            SELECT
                countMergeIf(aggCount_3, dateColumn > (now() - toIntervalDay(28))) AS s52,
                countMergeIf(aggCount_4, dateColumn > (now() - toIntervalDay(28))) AS s53
            FROM interpret_table_08 AS c
            PREWHERE cityHash64('0321352416546546546546546546546', lower('BU'), lower('random2'), toUInt128(joinGet(interpret_table_07, 'idColumnI64', toUInt128('1015029')))) = c.idColumnU64
        ) AS s08
    ) AS final_s01
    FORMAT JSONEachRow;
    