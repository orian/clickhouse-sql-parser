create materialized view mv_4 to mt_4 as
        select number as n, toString(number) as s from main_table where number % 13 != 4