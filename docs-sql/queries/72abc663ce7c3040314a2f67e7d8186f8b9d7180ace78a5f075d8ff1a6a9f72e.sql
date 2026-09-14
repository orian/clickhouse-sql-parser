-- Multi-name × multi-table in one ON-group (not a Cartesian product)
CREATE ROW POLICY p1, p2 ON t1, t2
    FOR SELECT USING id = 1
    TO accountant;

-- Different clusters per name in one statement
CREATE ROW POLICY pol1 ON CLUSTER cluster1 ON table1, pol2 ON CLUSTER cluster2 ON table2
