-- Aplicar la política de filas a todos los usuarios
CREATE ROW POLICY cdc_policy ON votes FOR SELECT USING _peerdb_is_deleted = 0 TO ALL;
