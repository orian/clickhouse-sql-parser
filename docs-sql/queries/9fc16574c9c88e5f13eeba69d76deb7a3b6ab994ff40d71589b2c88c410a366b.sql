       CREATE ROW POLICY estuary_columns ON system.columns FOR SELECT USING database = '<database>' TO <user>;
       CREATE ROW POLICY estuary_parts ON system.parts FOR SELECT USING database = '<database>' TO <user>;
       CREATE ROW POLICY estuary_tables ON system.tables FOR SELECT USING database = '<database>' TO <user>;
