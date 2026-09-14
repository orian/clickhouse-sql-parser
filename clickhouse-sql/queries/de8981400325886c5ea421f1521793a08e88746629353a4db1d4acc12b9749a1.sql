create materialized view mv_1 to mt_1 as
        select number as n, toString(number) as s from main_table where number % 13 != 0