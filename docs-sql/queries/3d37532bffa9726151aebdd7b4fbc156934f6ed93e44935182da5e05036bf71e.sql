clickhouse-cloud :) CREATE VIEW raw_data_parametrized AS SELECT * FROM raw_data WHERE id BETWEEN {id_from:UInt32} AND {id_to:UInt32}

CREATE VIEW raw_data_parametrized AS
SELECT *
FROM raw_data
WHERE (id >= {id_from:UInt32}) AND (id <= {id_to:UInt32})

Query id: 45fb83a6-aa55-4197-a7cd-9e1ad2c76d48

Ok.

0 rows in set. Elapsed: 0.102 sec.
