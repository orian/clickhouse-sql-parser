   ALTER TABLE website.clicks
   UPDATE visitor_id = getDict('visitors', 'new_visitor_id', visitor_id)
   WHERE visit_date < '2022-01-01'
