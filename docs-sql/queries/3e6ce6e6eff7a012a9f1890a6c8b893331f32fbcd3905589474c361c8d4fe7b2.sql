-- 创建角色
CREATE ROLE user_role

-- 授予 events 表只读权限。
GRANT SELECT ON default.events TO user_role
GRANT user_role TO user_1
GRANT user_role TO user_2
