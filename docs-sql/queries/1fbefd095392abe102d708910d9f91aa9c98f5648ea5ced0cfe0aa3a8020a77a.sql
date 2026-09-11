--mv1:begin
select a,b,c from {{ source('raw', 'table_1') }}
--mv1:end
union all
--mv2:begin
select a,b,c from {{ source('raw', 'table_2') }}
--mv2:end
