CREATE USER IF NOT EXISTS `pcm_scraper` IDENTIFIED WITH bcrypt_hash BY '<bcrypt-hash>';

GRANT SELECT ON `system`.`asynchronous_metric_log` TO `pcm_scraper`;
GRANT SELECT ON `system`.`metric_log` TO `pcm_scraper`;
GRANT SELECT ON `system`.`server_settings` TO `pcm_scraper`;
GRANT SELECT ON `system`.`tables` TO `pcm_scraper`;
GRANT SELECT ON `system`.`warnings` TO `pcm_scraper`;
GRANT SELECT ON `system`.`user_directories` TO `pcm_scraper`;
GRANT READ ON REMOTE TO `pcm_scraper`;
GRANT SYSTEM FLUSH LOGS ON *.* TO `pcm_scraper`;
