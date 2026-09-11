       -- 대상 데이터베이스 액세스: CREATE TABLE, DROP TABLE, SELECT, INSERT, TRUNCATE 등
       GRANT ALL ON <database>.* TO <user>;

       -- 메타데이터 디스커버리 및 파티션 관리를 위한 시스템 테이블 액세스
       -- 이는 위의 데이터베이스 권한에 포함되지 않습니다.
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
