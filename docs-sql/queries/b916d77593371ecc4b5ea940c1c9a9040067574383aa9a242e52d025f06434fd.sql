ALTER TABLE my_table MODIFY SETTING id_generator = 'sipHash64(tags)';
ALTER TABLE my_table MODIFY SETTING filter_by_min_time_and_max_time = 0;
ALTER TABLE my_table RESET SETTING filter_by_min_time_and_max_time;
