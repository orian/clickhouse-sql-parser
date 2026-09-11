CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- 使用安全的密码生成器

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
