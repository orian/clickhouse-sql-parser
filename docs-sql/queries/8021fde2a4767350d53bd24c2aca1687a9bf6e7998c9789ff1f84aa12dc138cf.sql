CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- use a secure password generator

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
