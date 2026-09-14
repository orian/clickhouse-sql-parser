-- ClickHouse Cloud にローカルテーブルの内容がすべて復元されていることを確認します
SELECT count() from nyc_taxi.trips_small_dist_local;
3000317
