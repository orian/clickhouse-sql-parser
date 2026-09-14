SELECT path, num_children, data_length
FROM system.keeper_storage
ORDER BY num_children DESC
LIMIT 3;
