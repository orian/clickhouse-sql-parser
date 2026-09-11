SELECT
    borough,
    count() AS complaints,
    bar(complaints, 0, 125000, 60)
FROM NYPD_Complaint
GROUP BY borough
ORDER BY complaints DESC
