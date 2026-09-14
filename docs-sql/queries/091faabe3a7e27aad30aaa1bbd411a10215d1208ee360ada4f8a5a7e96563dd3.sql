CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- use um gerador de senhas seguro

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
