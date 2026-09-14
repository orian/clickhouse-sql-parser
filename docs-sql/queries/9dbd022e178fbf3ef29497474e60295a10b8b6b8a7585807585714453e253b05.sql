-- Multiple names on one table target
CREATE [ROW] POLICY [IF NOT EXISTS | OR REPLACE] policy_name [, ...]
    [ON CLUSTER cluster_name]
    ON { [db.]table | db.* }
    [IN access_storage_type]
    [FOR SELECT] USING condition
    [AS {PERMISSIVE | RESTRICTIVE}]
    [TO {role1 [, role2 ...] | ALL | ALL EXCEPT role1 [, role2 ...]}]

-- One name on multiple table targets
CREATE [ROW] POLICY [IF NOT EXISTS | OR REPLACE] policy_name
    [ON CLUSTER cluster_name]
    ON { [db.]table | db.* } [, ...]
    [IN access_storage_type]
    [FOR SELECT] USING condition
    [AS {PERMISSIVE | RESTRICTIVE}]
    [TO {role1 [, role2 ...] | ALL | ALL EXCEPT role1 [, role2 ...]}]

-- Mixed packing: each name paired with its own table target
CREATE [ROW] POLICY [IF NOT EXISTS | OR REPLACE]
    policy_name ON { [db.]table | db.* } [, policy_name ON { [db.]table | db.* } ...]
    [ON CLUSTER cluster_name]
    [IN access_storage_type]
    [FOR SELECT] USING condition
    [AS {PERMISSIVE | RESTRICTIVE}]
    [TO {role1 [, role2 ...] | ALL | ALL EXCEPT role1 [, role2 ...]}]
