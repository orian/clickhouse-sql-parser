-- Origin SQL:
CREATE TABLE db.m (`id` UUID, `timestamp` DateTime64(3), `value` Float64) ENGINE = TimeSeries DATA db.m_data TAGS db.m_tags METRICS db.m_metrics;


-- Format SQL:
CREATE TABLE db.m (`id` UUID, `timestamp` DateTime64(3), `value` Float64) ENGINE = TimeSeries DATA db.m_data TAGS db.m_tags METRICS db.m_metrics;
