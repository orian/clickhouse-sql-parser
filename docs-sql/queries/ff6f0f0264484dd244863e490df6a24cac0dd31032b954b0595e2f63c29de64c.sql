➜  new ./clickhouse client --host myhost.us-central1.gcp.clickhouse.cloud --secure --user abcuser --ssh-key-file '/Users/testuser/.ssh/ch_key'
ClickHouse client version 23.12.1.863 (official build).
秘密鍵のパスフレーズを入力してください（パスフレーズなしの場合は空欄のままにしてください）:
ユーザー abcuser として myhost.us-central1.gcp.clickhouse.cloud:9440 に接続しています。
ClickHouse server バージョン 23.9.2 に接続しました。

clickhouse-cloud :) select currentUser();

SELECT currentUser()

Query id: d4b6bb60-ef45-47d3-8740-db9f2941dcd2

┌─currentUser()─┐
│ abcuser       │
└───────────────┘

1 row in set. Elapsed: 0.001 sec.

clickhouse-cloud :)
