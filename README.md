# Scalable Task Management API

This is a scalable task management API built with Go and the Fiber framework. The API supports advanced task tracking with role-based access, Redis caching, background workers, and CI/CD integration.

## Features

- **Task Management**: Create, update, and delete tasks.
- **Role-Based Access Control (RBAC)**: Admins, managers, and workers have different levels of access.
- **Redis Caching**: Cache frequently accessed task data for faster responses.
- **Background Workers**: Asynchronous task handling for things like sending reminders.
- **CI/CD Pipeline**: Automated testing and deployment.

## Tech Stack

- **Go**: The backend is built with Go and the Fiber framework.
- **PostgreSQL**: Used to store task data and user information.
- **Redis**: For caching data and background tasks.
- **JWT**: For role-based authentication and authorization.

## Setup

### Prerequisites

- Go 1.18+ installed
- PostgreSQL database set up
- Redis server running

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/task-management-api.git
   cd task-management-api
   ```
