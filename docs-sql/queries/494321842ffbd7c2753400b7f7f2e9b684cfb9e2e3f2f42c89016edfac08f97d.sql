-- Alterar o agendamento, removendo as configurações e dependências existentes.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE;

-- Alterar o agendamento e ajustar o comportamento de retentativas.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE
SETTINGS refresh_retries = 5,
         refresh_retry_initial_backoff_ms = 500,
         refresh_retry_max_backoff_ms = 60000;

-- Manter a dependência ao alterar o período.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR DEPENDS ON other_rmv;

-- Remover a dependência omitindo `DEPENDS ON`.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR;
