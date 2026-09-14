SELECT *, _part_offset FROM mergeTreeProjection(currentDatabase(), test, order_by_item_id);
