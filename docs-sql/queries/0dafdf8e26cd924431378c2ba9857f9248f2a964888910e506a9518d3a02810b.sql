SELECT currentUser(), authenticatedUser(); -- outputs "default    default"
CREATE USER james;
EXECUTE AS james SELECT currentUser(), authenticatedUser(); -- outputs "james    default"
