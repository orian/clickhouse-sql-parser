SET allow_experimental_ai_functions = 1;
SELECT aiRedact('Purchase was done by customer John Doe with email test@test.org', ['email', 'credit_card', 'name'])
