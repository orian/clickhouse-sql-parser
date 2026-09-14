CREATE TABLE my_table
ENGINE = TimeSeries
SETTINGS tags_to_columns = {'instance': 'instance', 'job': 'job'}
