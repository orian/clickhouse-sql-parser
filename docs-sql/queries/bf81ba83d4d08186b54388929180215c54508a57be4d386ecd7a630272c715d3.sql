SELECT * FROM redis_table JOIN merge_tree_table ON merge_tree_table.key=redis_table.key;
