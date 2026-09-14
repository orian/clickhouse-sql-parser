  SELECT count()
  FROM file('primary-hits_UserID_URL.idx', 'RowBinary', 'UserID UInt32, URL String');
