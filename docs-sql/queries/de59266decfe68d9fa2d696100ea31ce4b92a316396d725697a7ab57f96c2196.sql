-- 查看过去一小时内活跃度最高的前 10 名用户
SELECT user\_id, COUNT(*) AS actionsFROM analytics.eventsWHERE event\_time >= now() - INTERVAL 1 HOURGROUP BY user\_idORDER BY actions DESCLIMIT 10;
