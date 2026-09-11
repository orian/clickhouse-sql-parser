   ALTER TABLE website.clicks
   UPDATE url = substring(url, position(url, '://') + 3), visitor_id = new_visit_id
   WHERE visit_date < '2022-01-01'
