select one from (select main,sum(two) as one from tabl where x>1) where main like '%olsztyn%' order by one desc limit 10
