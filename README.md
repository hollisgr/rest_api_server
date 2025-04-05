### START SERVER

- To start server use make;

### ENDPOINTS

- GET /users -- list of users -- 200, 404, 500
- GET /users/:id -- user by id -- 200, 404, 500
- POST /users/:id -- create user -- 200, 4xx, Header Location: url
- DELETE /users/:id -- delete user by id -- 204, 404, 400

### TEST SERVER

- Use make users;
- Use make user;
- Use make create;
- Use make delete;
