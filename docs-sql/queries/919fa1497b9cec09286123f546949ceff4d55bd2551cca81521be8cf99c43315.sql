  CREATE USER userID IDENTIFIED WITH sha256_hash by 'hashed_password';
  GRANT default_role to userID;
