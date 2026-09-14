CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- используйте надёжный генератор паролей

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
