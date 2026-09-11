-- Several policy names, one table
CREATE ROW POLICY pol1, pol2, pol3 ON table1
    FOR SELECT USING id = 1
    TO accountant;

-- One policy name, several tables
CREATE ROW POLICY IF NOT EXISTS pol1 ON table1, table2, table3
    FOR SELECT USING id = 1
    TO accountant;

-- Mixed packing: different name per table
CREATE ROW POLICY p4 ON db.table, p5 ON db2.table2
    USING a = b;

-- Same policy on several tables, on a cluster
CREATE ROW POLICY IF NOT EXISTS pol1 ON CLUSTER replicated_cluster ON table1, table2
    FOR SELECT USING id = 1
    TO accountant;
