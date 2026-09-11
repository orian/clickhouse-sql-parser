   ALTER TABLE clicks ON CLUSTER main_cluster
   UPDATE click_count = click_count / 2
   WHERE visitor_id ILIKE '%robot%'
