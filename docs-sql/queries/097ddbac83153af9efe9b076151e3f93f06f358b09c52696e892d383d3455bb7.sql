-- ロールを作成する
CREATE ROLE user_role

-- eventsテーブルへの読み取り専用権限を付与する。
GRANT SELECT ON default.events TO user_role
GRANT user_role TO user_1
GRANT user_role TO user_2
