-- Crear rol
CREATE ROLE user_role

-- Otorgar acceso de solo lectura a la tabla events.
GRANT SELECT ON default.events TO user_role
GRANT user_role TO user_1
GRANT user_role TO user_2
