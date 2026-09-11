SYSTEM ENABLE FAILPOINT replicated_merge_tree_insert_retry_pause;
SELECT * FROM system.fail_points WHERE enabled = 1
