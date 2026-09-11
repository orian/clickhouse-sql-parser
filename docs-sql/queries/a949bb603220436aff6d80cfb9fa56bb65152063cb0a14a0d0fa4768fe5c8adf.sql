-- Criar role
CREATE ROLE user_role

-- Conceder acesso somente leitura à tabela events.
GRANT SELECT ON default.events TO user_role
GRANT user_role TO user_1
GRANT user_role TO user_2
