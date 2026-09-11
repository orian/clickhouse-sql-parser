WITH
   Region = 'us' AS segment1,
   Browser = 'Chrome' AS segment2
SELECT
   uniqIf(UserID, segment1),
   uniqIf(UserID, segment2)
WHERE segment1 OR segment2
