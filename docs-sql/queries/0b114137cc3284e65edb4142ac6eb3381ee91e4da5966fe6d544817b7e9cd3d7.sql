-- Rename: exactly one fully qualified policy (one name on one target).
-- RENAME TO may be combined with the same optional alteration clauses as below.
ALTER [ROW] POLICY [IF EXISTS] name
    ON { [database.]table | database.* }
    RENAME TO new_name
    [ON CLUSTER cluster_name]
    [AS {PERMISSIVE | RESTRICTIVE}]
    [FOR SELECT]
    [USING {condition | NONE}][,...]
    [TO {role [,...] | ALL | ALL EXCEPT role [,...]}]

-- Multiple names on one table target (no RENAME)
ALTER [ROW] POLICY [IF EXISTS] name [, ...]
    [ON CLUSTER cluster_name]
    ON { [database.]table | database.* }
    [AS {PERMISSIVE | RESTRICTIVE}]
    [FOR SELECT]
    [USING {condition | NONE}][,...]
    [TO {role [,...] | ALL | ALL EXCEPT role [,...]}]

-- One name on multiple table targets (no RENAME)
ALTER [ROW] POLICY [IF EXISTS] name
    [ON CLUSTER cluster_name]
    ON { [database.]table | database.* } [, ...]
    [AS {PERMISSIVE | RESTRICTIVE}]
    [FOR SELECT]
    [USING {condition | NONE}][,...]
    [TO {role [,...] | ALL | ALL EXCEPT role [,...]}]

-- Mixed packing: each name paired with its own table target (no RENAME)
ALTER [ROW] POLICY [IF EXISTS]
    name ON { [database.]table | database.* } [, name ON { [database.]table | database.* } ...]
    [ON CLUSTER cluster_name]
    [AS {PERMISSIVE | RESTRICTIVE}]
    [FOR SELECT]
    [USING {condition | NONE}][,...]
    [TO {role [,...] | ALL | ALL EXCEPT role [,...]}]
