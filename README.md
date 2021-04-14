# Kanban Board


## Project Directory

```bash
.
├── Backend
│   ├── Auth
│   │   ├── Auth.go
│   │   ├── db
│   │   │   ├── setup.go
│   │   │   └── User.go
│   │   ├── helper
│   │   │   ├── CheckHash.go
│   │   │   ├── CheckIfUserExists.go
│   │   │   ├── CompareUsers.go
│   │   │   ├── CreateJWT.go
│   │   │   ├── hashPassword.go
│   │   │   └── ValidateToken.go
│   │   ├── user.db
│   │   └── views
│   │       ├── Demo.go
│   │       ├── LoginViews.go
│   │       └── RegisterViews.go
│   ├── Backend
│   ├── go.mod
│   ├── go.sum
│   ├── Main
│   │   ├── Backend
│   │   ├── controllers
│   │   │   └── todo.go
│   │   ├── favicon.ico
│   │   ├── Main.go
│   │   ├── middleware
│   │   │   ├── GetUser.go
│   │   │   └── TokenAuthMiddleWare.go
│   │   ├── models
│   │   │   ├── setup.go
│   │   │   └── todo.go
│   │   └── test.db
│   ├── main.go
│   ├── test.db
│   ├── user.db
│   └── users.db
├── Design
│   ├── css
│   │   ├── login.css
│   │   └── main.css
│   ├── index.html
│   ├── js
│   │   ├── index.js
│   │   ├── load_localStorage.js
│   │   └── login.js
│   └── login.html
└── README.md

```

Backend is where all our Go Files are.
We have 2 apps Auth and Main. Auth is used to Authenticate and Authorize Users and Main is used to connect to our Database.

Design is Frontend Part of this App.
We have static files like css and js files in seprate folders.
`index.html` is the entry-point for our app.


## Tech Stack

### For Backend - 

- Go/Gin Server
- JWT's for Authentication
- SQLite3 for Database (Migrating to PostgreSQL)

### For Frontend - 

- HTML / CSS
- Javascript/JQuery
- Ionic Framework for CSS and Icons


## Main Purpose Of this App 

We can say that this App is an **Trello Clone**, But it was made out of my need for good Kanban Board Web App.

Kanban Board is a Board where we can Create our own **columns** according to our need and Seprate & Manage tasks


## Learnings Building this App

I have never used Go for building Web Server before, this was my first time and So Far, I have learnt a lot.

Before I was using Django for my Personal Projects but it is "Batteries Included Framework", So there was no need for implementing some things like Authentication, Authorization. I learnt have to use JWT's for a secure Login and how to create middlewares for not letting non registered users have access.

This was also my First Time Developing Backend & Frontend Seperatly. So, I learned about Nginx and started using it. I have never used some like that before.

This Project is Blessing in Disguise as I learnt a lot basic things of Full Stack Development which I wouldn't have learnt if I were to use Django.

*No Offence To Django Users, I am a fan of Django and I know I can implement this things in Django but the "Batteries" were too good for me to stop using them as a beginner*


## Contributing

- Fork this Project
- Clone forked Project
- Create your branch
- Run this command in Backend/
  ```go get -u ./...```
- Run this Command in Backend/
  ```go run server.go``` OR ```./Backend```
- Start a Server in Design Server to Serve Files on Port 8000
- Run an Nginx Server with `/api` pointing to Gin Server and `/` to Design Server

Open `http://localhost` and you are good to go


## Future of this Project

- I am planning to improve Frontend, (Use some JS Framework?)
- Improving Security of App
	- Currently, I am storing JWT's in Local Storage But I am planning to shift to Redis
- Implementing More Ideas