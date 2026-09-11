SELECT grants.user_name,
grants.role_name,
users.name AS role_member,
grants.access_type,
grants.database,
grants.table
FROM system.grants LEFT OUTER JOIN system.role_grants ON grants.role_name = role_grants.granted_role_name
LEFT OUTER JOIN system.users ON role_grants.user_name = users.name

UNION ALL

SELECT grants.user_name,
grants.role_name,
role_grants.role_name AS role_member,
grants.access_type,
grants.database,
grants.table
FROM system.role_grants LEFT OUTER JOIN system.grants ON role_grants.granted_role_name = grants.role_name
WHERE role_grants.user_name is null;
