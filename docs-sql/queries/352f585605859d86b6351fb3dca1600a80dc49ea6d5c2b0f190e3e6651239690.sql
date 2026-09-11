INSERT INTO tw_weather_data SELECT *
FROM url('https://storage.googleapis.com/taiwan-weather-observaiton-datasets/daily_weather_preprocessed_1896_2023.csv', 'CSVWithNames')

