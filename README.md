# Ticketing + Gamification Backend

This is a RESTful API backend built with **Go (Golang)** and **Fiber v2** for the ticketing platform. It uses **PostgreSQL** as the database (managed via **GORM**) and is fully containerized with **Docker** for easy "plug-and-play" development.

## Requirements

- **Docker Desktop** (Must be running)
- Git

## Getting started

> **Note:** This project uses **Vendor Mode**. You **DO NOT** need to install Go or download dependencies manually. All libraries are included in the `vendor/` folder for fast, offline-ready builds.

### 1. Clone the repository
Clone this project to your local machine.

### 2. Configure Environment Variables
Create a new file named `.env` in the root directory. 

### 3. Start the Application
Run the following command in your terminal:

```bash
docker-compose up --build
```

The server will start at `http://localhost:8080`.

---

## Database Access

To connect to the database using tools like **DBeaver** or **PgAdmin**, use the credentials you defined in your `.env` file:

- **Host:** `localhost`
- **Port:** `5432`
- **Database:** (value of `DB_NAME`)
- **Username:** (value of `DB_USER`)
- **Password:** (value of `DB_PASSWORD`)

---

## Available Commands

Since we use Docker, here are the common commands you will need:

| Command | Description |
| :--- | :--- |
| `docker-compose up --build` | **Start the API and DB.** Recompiles code changes (use this if you edited Go code). |
| `docker-compose up` | Start without rebuilding (faster, if no code changed). |
| `docker-compose down` | Stop all containers. |
| `docker-compose down -v` | **Hard Reset.** Stops containers AND deletes the database data (Volume). Use this if you want to wipe the DB clean or if you changed the DB password in .env. |

---

## Architecture Overview

- **Go 1.25 + Fiber v2:** High-performance web framework.
- **PostgreSQL + GORM:** Robust relational database with Object-Relational Mapping.
- **Docker + Vendor Mode:** Ensures the environment is identical on every machine and works even with unstable internet.
- **JWT Authentication:** Secure stateless authentication using `golang-jwt/v5`.
- **Environment Variables:** Sensitive data (passwords, secrets) are loaded from `.env` and never hardcoded in Docker Compose.

### Project Structure
- `models/`: Database schemas (User, Project, Ticket, Gamification).
- `controllers/`: Business logic and request handling.
- `routes/`: API endpoint definitions.
- `middleware/`: JWT protection and request validation.

---

## Implemented Features (API)

Currently, the backend supports the following core features:

### 1. Authentication
- `POST /api/register` - Register a new user (Auto-creates UserStats).
- `POST /api/login` - Login and receive **Bearer Token** (JWT).

### 2. Projects (Protected)
- `POST /api/projects` - Create a new project (Creator automatically becomes the Owner).
- `GET /api/projects` - Get list of projects the user belongs to.

### 3. Ticketing (Protected)
- `POST /api/tickets` - Create a ticket (Task/Bug/Feature) inside a project.
- `GET /api/tickets?project_id=1` - Get all tickets for a specific project.

---

## How to Update Dependencies

If you need to add a new library (e.g., payment gateway), follow these steps to keep the **Vendor Mode** working:

1. Run `go get <package_name>` locally in your terminal.
2. Run `go mod vendor` to update the local vendor folder.
3. **Push the changes** (including the updated `vendor/` folder) to Git.
4. Teammates just need to `git pull` and `docker-compose up --build`.
