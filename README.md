# API Endpoints

## Authentication

- **POST /memberships/sign-up**  
  Register a new user.

- **POST /memberships/login**  
  Login a user.

- **POST /memberships/logout**  
  Logout the user.

- **POST /memberships/refresh-token**  
  Refresh the JWT token / Refresh access token.

## Posts

- **GET /posts**  
  Get a paginated list of posts.

- **POST /posts/create**  
  Create a new post.

- **GET /posts/:postID**  
  Get post details.

- **POST /posts/comment/:postID**  
  Add a comment to a post.

- **PUT /posts/user_activity/:postID**  
  Like/unlike a post.

## Configuration

The application uses a YAML configuration file located at `internal/configs/config.yaml`. Configurable settings include:

- **Service port**
- **Database connection**
- **JWT secret key**
- **Logging settings**

### Example Configuration:

```yaml
service:
  port: ":9876"
  secretJWT: "secretpassword"

database:
  dataSourceName: "<username>:<password>@tcp(<host>:<port>)/<dbname>?parseTime=true"
```
# Development
## Prerequisites
Go 1.23.1 or higher
MySQL
Make (for running migration commands)
Database Migrations

Create new migration:

```yaml
make migrate-create name=migration_name
```

Apply migrations:
```yaml
make migrate-up
```

Rollback migrations:
```yaml
make migrate-down
```

## Project Architecture
The project follows a clean architecture pattern with the following layers:

- Handlers (HTTP request handling)
- Services (Business logic)
- Repositories (Data access)
- Models (Data structures)


## Security
- JWT-based authentication
- Password hashing using bcrypt
- Refresh token mechanism for session management
- Middleware for protected routes
