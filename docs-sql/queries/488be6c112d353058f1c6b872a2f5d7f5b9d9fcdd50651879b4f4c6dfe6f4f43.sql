create table db1.table1_exchange_temp
(
 id Int32,
 string_field String
)
engine = MergeTree()
order by id;
