-- Cambiar el horario, eliminar la configuración y las dependencias existentes.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE;

-- Cambiar el horario y ajustar el comportamiento de reintentos.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE
SETTINGS refresh_retries = 5,
         refresh_retry_initial_backoff_ms = 500,
         refresh_retry_max_backoff_ms = 60000;

-- Mantener la dependencia al cambiar el período.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR DEPENDS ON other_rmv;

-- Eliminar la dependencia omitiendo `DEPENDS ON`.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR;
