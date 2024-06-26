
# gocool

This project is a comprehensive Go framework designed for efficient web service development, specifically focusing on user management. It includes functionality for user registration, authentication, and profile management, among other features.

## Setup

To run this project, you need to have Docker installed. You can build and run the container using the following commands:

```bash
docker build -t gocool .
docker run -p 8080:8080 gocool
```

## Configuration

Configure the application settings in the `.env` file according to your environment setup.

## Features

- **User Registration**: Allows new users to register.
- **User Authentication**: Supports user authentication using JWT.
- **User Profile Management**: Users can update their profile details.
- **Data Caching**: Implements caching for user data to enhance performance.

Each component is modular and well-integrated, ensuring a seamless development experience and easy maintenance.
