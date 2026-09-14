    -- スキーマ情報を読み取る権限を付与
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- データベースとテーブルへの読み取りアクセスを付与
    GRANT SELECT ON your_database.* TO your_databrain_user;
