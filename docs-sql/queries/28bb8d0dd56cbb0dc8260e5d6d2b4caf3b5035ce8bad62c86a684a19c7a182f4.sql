-- 行ポリシーを作成する
CREATE ROW POLICY user_filter_1 ON default.events USING tenant_id=1 TO user_1
CREATE ROW POLICY user_filter_2 ON default.events USING tenant_id=2 TO user_2
