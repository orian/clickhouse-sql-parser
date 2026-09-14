CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- 安全なパスワードジェネレーターを使用すること

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
