CREATE ROW POLICY policy_name ON t1
FOR SELECT USING _peerdb_is_deleted = 0;
