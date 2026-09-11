SELECT table_name, 
       (xpath('/row/cnt/text()', xml_count))[1]::text::bigint AS row_count
FROM (
  SELECT table_name, 
         query_to_xml('SELECT count(*) AS cnt FROM ' || quote_ident(table_name), false, true, '') AS xml_count
  FROM information_schema.tables
  WHERE table_schema = 'public'
) t
ORDER BY table_name;
