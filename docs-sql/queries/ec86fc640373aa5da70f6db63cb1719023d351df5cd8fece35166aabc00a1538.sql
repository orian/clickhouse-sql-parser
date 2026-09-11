CREATE USER fivetran_user IDENTIFIED BY '<password>'; -- utiliza un generador de contraseñas seguro

GRANT CURRENT GRANTS ON *.* TO fivetran_user;
