INSERT INTO dict SELECT rand(), rand()%1000, rand()*0.0001, toString(number),
        toString(number), toString(number), toString(number), toString(number), toString(number),
        toString(number), toString(number) FROM numbers(1000000);