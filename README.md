# go-gin-api-starter

English | [中文](https://github.com/hunter-ji/go-gin-api-starter/blob/main/doc/README.zh-CN.md)

## Project Introduction

go-gin-api-starter is a RESTful API project starter based on the Gin framework. It helps developers quickly build and
start scalable Go backend services. This starter template focusing on API development. It provides a clear and easy project base.

You can also check out other templates I've put together:

- Frontend template: [vue-ts-tailwind-vite-starter](https://github.com/hunter-ji/vue-ts-tailwind-vite-starter)
- Database template: [postgres-redis-dev-docker-compose](https://github.com/hunter-ji/postgres-redis-dev-docker-compose)

## Project Features

- **Gin Framework**: Uses Gin's high performance and flexibility to quickly build RESTful APIs.
- **MVC**: Uses MVC design to make code easier to maintain and test.
- **PostgreSQL Support**: Uses PostgreSQL as the main database for reliable data and strong query abilities.
- **Redis Support**: Uses Redis for caching management to improve app performance.
- **Modular Structure**: Clear directory structure for easy project expansion and maintenance.
- **Middleware Support**: Includes common middlewares like logging, error handling, and authentication.
- **Config Management**: Flexible config management, supports multi-environment deployment.
- **Custom Validators**: Built-in custom format validators, registered in Gin's binding for easy expansion.
- **JWT Authentication**: Including auto token refresh and active refresh requests, easy to use.
- **User Module**: Pre-set user module and test code for quick development reference and base.
- **API Version Control**: Built-in API version control for managing different API versions.
- **Enhanced Dev Mode Logging**: In dev mode, automatically prints detailed HTTP request and response info, greatly
  helping API debugging and development.

## Quick Start

Follow these steps to quickly set up and run your project:

### 1. Clone the Project

```bash
git clone https://github.com/hunter-ji/go-gin-api-starter
cd go-gin-api-starter
```

### 2. Initialize Project

Use the provided `project_bootstrap.sh` script to quickly initialize your project:

```bash
bash project_bootstrap.sh
```

The process is as follows:

```
🚀 Welcome to the Go Project Initializer!

📝 Please enter the new project name:
> new-project

Initializing your project...
Project name updated successfully!

🐰 Do you need RabbitMQ in your project? [Y/n]:
> n
RabbitMQ folder removed.

🗑️ Do you want to remove the existing .git directory? [Y/n]:
>
Existing .git directory removed.

🔧 Do you want to initialize a new Git repository? [Y/n]:
>
Initialized empty Git repository in /path/to/new-project/.git/
New Git repository initialized successfully.

🎉 Project initialization complete!
Summary:
   Old project name: Template
   New project name: new-project
   RabbitMQ: Removed
   Git: Reinitialized

🧹 Do you want to run 'go mod tidy' to clean up dependencies? [Y/n]:
> n
Skipped 'go mod tidy'. Remember to run it later if needed.

🚀 How to start your project:
1. Modify .env.development in the project root to set Redis and PostgreSQL configurations.
2. Run the following command to start your project:
   make run

🎈 Your project is ready! Happy coding!
```

This will help you:

- Change project name
- Remove unneeded modules, like RabbitMQ
- Delete/reinitialize .git folder
- Install dependencies

### 3. Configure Database

Add Redis and PostgreSQL configs in the .env.development file. Example config:

```
# postgresql
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_DATABASE_NAME=

# redis
REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=
REDIS_DB=0
```

For a quick dev environment database, you can use
my [Database Docker Compose](https://github.com/hunter-ji/postgres-redis-dev-docker-compose).

### 4. Start Project

After config, use this command to start the project:

```bash
make run
```

### 5. Verify

Open a browser or use curl to access:

```
http://localhost:9000/api/ping
```

If all is well, you should see a "pong" response.

## Project Structure

### `/cmd`

Main project applications.

- `/api`: Entry point for API server.

### `/config`

Config files and loading logic.

- config.go: Main config file.
- constant.go: Constant definitions.
- environment_variable.go: Environment variable handling.
- router_white_list.go: Router whitelist config.

### `/doc`

Design docs, user docs, and other project-related docs.

### `/internal`

Private app and library code.

- `/api`: API layer, defines routes and handler functions.
- `/constant`: Internal constants.
- `/database`: Database connection and init logic.
- `/middleware`: HTTP middleware.
- `/model`: Data models and struct definitions.
- `/repository`: Data access layer, handles data persistence.
- `/service`: Business logic layer.

### `/migration`

Database migration files.

- `000001_init_schema.down.sql`: Rollback script for initial schema.
- `000001_init_schema.up.sql`: Execution script for initial schema.

### `/pkg`

Library code that can be used by external apps.

- /auth: Auth-related functions.
- /util: Common utility functions.

### `/test`

Additional external test apps and test data.

- `ping_test.go`: Ping function test.
- `user_test.go`: User-related function tests.

### Root Directory Files

- `Makefile`: Makefile defining common commands.
- `project_bootstrap.sh`: Project init script.

## Running Tests

This project offers various ways to run tests for different testing needs.

### Run All Tests

Run all tests in the project:

```bash
make test
```

### Run Tests in a Specific Folder

Run tests in the ./test folder:

```bash
make test-folder
```

### Run a Specific Test File

Run a specified test file:

```bash
make test-file FILE=path/to/your_test.go
```

For example:

```bash
make test-file FILE=./test/user_test.go
```

### Output Detailed Test Logs

For all the above test commands, you can add -verbose, for example:

```bash
make test-verbose
make test-folder-verbose
make test-file-verbose FILE=./test/user_test.go 
```

## Switching Run Environments

The project relies on the `NODE_ENV` environment variable to switch environments. If not set, it defaults to
`development`.
The project loads different config files based on the environment, like `.env.development`, `.env.test`,
`.env.production`, etc.
Just create a new `.env` file and set the `NODE_ENV` variable.

## More Usage Methods

For more usage methods, you can refer to:

- The `Makefile` file, which defines common commands like run, test, build, clean, etc.
- README files inside module folders
- Code comments

The documentation will be continuously updated and improved in the future.

## Frequently Asked Questions (FAQ)

### Q1: How to change the default server port?

A: Modify the `SERVER_PORT` variable in the `.env.development` file. For example, setting it to `SERVER_PORT=8080` will
make
the server run on port 8080.

### Q2: Which databases does the project support?

A: Currently, the project mainly supports PostgreSQL as the main database. If you need to use other databases, you may
need to modify the database connection code in `internal/database`.

### Q3: How to add new API routes?

A: Add new handler functions in the appropriate version folder (like v1) under the `internal/api` directory, then
register
this new route in the `internal/api/router.go` file.

### Q4: How to customize middleware?

A: Create a new middleware file in the `internal/middleware` directory, implement the middleware logic, then use the
`r.Use()` method in `internal/api/router.go` to apply this middleware.

### Q5: How to turn on or off the enhanced logging feature in development mode?

A: This feature is automatically enabled in development mode by default. If you want to disable it in production, make
sure the `NODE_ENV` environment variable is not set to `development`.

### Q6: How to run specific tests?

A: Use the make test-file command, specifying the test file to run. For example:

```bash
make test-file FILE=./test/user_test.go
```

### Q7: How to configure JWT token expiration time?

A: Modify variables like `accessTokenExpire` in the `pkg/auth/token.go` file.

### Q8: How to add new custom validators?

A: Add new validation functions in the `pkg/util/formatValidator` directory, then register this new validator in
`pkg/util/customBindValidator`.

### Q9: Does the project support CORS?

A: Yes, the project has CORS middleware configured by default.

## Author's Note

In recent years when writing Go projects, I summarized a template for myself. When developing new projects recently, I
found some shortcomings and areas that were not fully improved before, which I hadn't systematically updated and
improved. This time, before developing the project, I started from scratch and reorganized it. Not only did it benefit
myself, but I also hope it can help other developers.

This project will have some of my personal style, such as directory structure, naming conventions, etc. Of course, these
can be adjusted according to your own preferences and project requirements.

If you have any suggestions, you are welcome to issue or PR.

