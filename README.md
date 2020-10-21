# todo-svr 
## gloang code for assigning tasks to user


### endpoints:
1. /task

   GET 	  : get all tasks in database (dev) [x]
   GET   /id? : get tasks correspondingto your ID (user) [x]

   POST 	: make a task to yourself. [x]
   POST /id : make a task to your assignee id [x]

   PUT /id : change tasks based on its id (user) []
   DELETE /id : delete tasks based on its id ( user) [] 

2. /user

	GET : get all users in database (dev) [x]
	POST /register : register a user in db [x]
	DELETE /id : delete user with id ( dev) []
	PUT /id : update user with id []

3. /login
	GET: show login page [x]
	POST: login as sb with specific role (dev,user) [x]

4. POST /supervise/id : make user with (id) your assignee (with his consent) []
5. POST /allow/id : make user with (id) your admin (with his consent) []


### roles: 
- dev 
- user

### What hasn't been yet to implemented?
- An adequate testing scheme for every compilation.
- replace lengthy code ( check error).
- A proper frontend 
- ElasticSearch + Redis ( stolen ideas)
- TLS setup?
