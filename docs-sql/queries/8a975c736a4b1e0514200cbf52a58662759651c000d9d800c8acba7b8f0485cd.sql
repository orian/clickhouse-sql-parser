-- 역할 생성
CREATE ROLE user_role

-- events 테이블에 읽기 전용 권한 부여.
GRANT SELECT ON default.events TO user_role
GRANT user_role TO user_1
GRANT user_role TO user_2
