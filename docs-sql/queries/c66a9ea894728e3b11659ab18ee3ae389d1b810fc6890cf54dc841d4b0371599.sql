    -- 스키마 정보를 읽기 위한 권한 부여
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- 데이터베이스 및 테이블에 대한 읽기 권한 부여
    GRANT SELECT ON your_database.* TO your_databrain_user;
