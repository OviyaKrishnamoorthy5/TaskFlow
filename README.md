# TaskFlow 🚀

TaskFlow is a full-stack task management application currently under development.

## 🛠️ Tech Stack

### Frontend
- React
- TypeScript
- Tailwind CSS

### Backend
- Go
- Gin
- REST APIs
- JWT Authentication
- bcrypt

### Database
- PostgreSQL

## ✨ Features

- User Registration
- User Login
- JWT Authentication
- User Profile
- Create Tasks
- Assign Tasks
- Update Tasks
- Delete Tasks
- Search Tasks
- Filter Tasks
- Pagination
- Responsive Dashboard

## 📁 Project Structure

TaskFlow/
├── Backend/
│   ├── config/
│   ├── controllers/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   ├── utils/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
│
└── Frontend/

## 🔐 Authentication

TaskFlow uses JWT-based authentication to secure protected API endpoints.

User passwords are securely hashed using bcrypt before being stored in the database.

## 🚧 Project Status

**Work in Progress**

### Completed

- Go backend setup
- Gin REST API
- PostgreSQL database connection
- User registration API
- User login API
- JWT authentication
- Protected profile endpoint

### In Progress

- Task management APIs
- React frontend
- Dashboard
- Task search and filtering
- Pagination
- Frontend and backend integration

📌 API Endpoints
Method	Endpoint	Description
GET	/ping	Check backend status
POST	/register	Register a new user
POST	/login	Login
GET	/profile	Get user profile
🔮 Future Improvements
Role-based access control
Task priorities and deadlines
Unit and integration testing
Logging and monitoring
Docker support
CI/CD pipeline

TaskFlow is an ongoing project that will be continuously updated with new features.
