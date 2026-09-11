-- See top 10 active users in the last hour
SELECT user\_id, COUNT(*) AS actionsFROM analytics.eventsWHERE event\_time >= now() - INTERVAL 1 HOURGROUP BY user\_idORDER BY actions DESCLIMIT 10;
