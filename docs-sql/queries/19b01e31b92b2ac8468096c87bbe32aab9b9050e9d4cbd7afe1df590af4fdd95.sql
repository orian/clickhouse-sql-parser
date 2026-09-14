CREATE TABLE my_table ENGINE=TimeSeries
SETTINGS id_generator = 'sipHash64(tags)'
