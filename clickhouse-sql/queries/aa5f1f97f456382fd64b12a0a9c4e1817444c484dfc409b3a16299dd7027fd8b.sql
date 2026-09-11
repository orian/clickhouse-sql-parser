create materialized view mv_3 to mt_3 as
        select number as n, toString(number) as s from main_table where number % 13 != 3