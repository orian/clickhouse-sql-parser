
        ALTER TABLE projection_materialize ADD PROJECTION agg_proj (SELECT key % 1000 AS k, sum(val) GROUP BY k)
    