create materialized view mv_2 to mt_2 as
        select number as n, toString(number) as s from main_table where number % 13 != 1