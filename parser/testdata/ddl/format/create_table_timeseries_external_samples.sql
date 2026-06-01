-- Origin SQL:
CREATE TABLE db.m (`id` UUID, `timestamp` DateTime64(3), `value` Float64) ENGINE = TimeSeries SAMPLES db.m_samples TAGS db.m_tags METRICS db.m_metrics;


-- Format SQL:
CREATE TABLE db.m (`id` UUID, `timestamp` DateTime64(3), `value` Float64) ENGINE = TimeSeries SAMPLES db.m_samples TAGS db.m_tags METRICS db.m_metrics;
