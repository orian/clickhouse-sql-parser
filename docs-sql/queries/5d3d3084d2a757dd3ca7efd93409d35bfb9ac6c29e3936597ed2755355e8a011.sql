       -- 移行先データベースへのアクセス: CREATE TABLE、DROP TABLE、SELECT、INSERT、TRUNCATE など。
       GRANT ALL ON <database>.* TO <user>;

       -- メタデータ検出およびパーティション管理のためのシステムテーブルへのアクセス。
       -- これらは上記のデータベース権限には含まれません。
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
