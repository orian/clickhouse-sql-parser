SYSTEM RESTORE REPLICA table_name;  -- すべてのパーティションをデタッチし、ZK 内のメタデータを再作成（新しい空のテーブルとして扱われます）した後、すべてのパーティションを再アタッチします

SYSTEM SYNC REPLICA table_name; -- レプリカがパーツを同期するまで待機します。また、リカバリ完了後はすべてのレプリカで `system.detached_parts` を確認することを推奨します。
