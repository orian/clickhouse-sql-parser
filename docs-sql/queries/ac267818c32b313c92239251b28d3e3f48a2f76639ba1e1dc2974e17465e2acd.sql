CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- 안전한 비밀번호 생성기를 사용하십시오

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
