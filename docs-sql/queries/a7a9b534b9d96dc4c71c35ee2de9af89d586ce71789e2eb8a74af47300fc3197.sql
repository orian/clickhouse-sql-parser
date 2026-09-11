CREATE USER u1;
EXECUTE AS u1 SELECT currentUser(), authenticatedUser();
DROP USER u1;
