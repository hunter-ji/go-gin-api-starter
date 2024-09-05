# go-gin-api-starter

[English](https://github.com/hunter-ji/go-gin-api-starter) | 中文

## 项目介绍

go-gin-api-starter是一个基于Gin框架的RESTful API项目模板，旨在帮助开发者快速构建和启动高效、可扩展的Go后端服务。这个模板采用了MVC架构的变体，专注于API开发，提供了一个结构清晰、易于扩展的项目基础。

如果你有兴趣，也可以看看我整理的别的模板：

- 前端模板：[vue-ts-tailwind-vite-starter](https://github.com/hunter-ji/vue-ts-tailwind-vite-starter)
- 数据库模板：[postgres-redis-dev-docker-compose](https://github.com/hunter-ji/postgres-redis-dev-docker-compose)

## 项目特性

- **Gin 框架**: 利用Gin的高性能和灵活性，快速构建RESTful API。
- **MVC 架构**: 采用MVC的设计理念，实现关注点分离，提高代码的可维护性和可测试性。
- **模块化结构**: 清晰的目录结构，便于项目的扩展和维护。
- **中间件支持**: 包含常用中间件，如日志记录、错误处理和认证等。
- **配置管理**: 灵活的配置管理，支持多环境部署。
- **自定义验证器**: 内置自定义的格式化验证器，已注册到Gin的binding中，可灵活扩展。
- **JWT 认证**: 封装了JWT认证机制，包括token自动刷新和主动请求刷新功能，可直接使用。
- **开发模式日志增强**: 在开发模式下，自动打印详细的HTTP请求和响应信息，极大方便了API的调试和开发过程。
- **PostgreSQL 集成**: 使用PostgreSQL作为主数据库，确保数据的可靠性和强大的查询能力。
- **Redis 支持**: 集成Redis用于缓存和会话管理，提升应用性能。
- **用户模块**: 预置用户模块及其测试代码，为快速开发提供参考和基础。
- **API 版本控制**: 内置API版本控制机制，便于管理不同版本的API。

## 快速开始

按照以下步骤快速设置和运行您的项目：

### 1. 克隆项目

```bash
git clone https://github.com/hunter-ji/go-gin-api-starter
cd go-gin-api-starter
```

### 2. 初始化项目

使用提供的 `project_bootstrap.sh` 脚本来快速初始化您的项目：

```bash
bash project_bootstrap.sh
```

过程如下：

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

这将会帮助你完成以下操作：

- 更换项目名称
- 删除不需要的模块，比如`RabbitMQ`
- 删除/重新初始化`.git`文件夹
- 安装依赖

### 3. 配置数据库

在 `.env.development` 文件中添加 Redis 和 PostgreSQL 的配置。示例配置如下：

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

若需要快速启动开发环境的数据库，可以使用我的 [Database Docker Compose](https://github.com/hunter-ji/postgres-redis-dev-docker-compose) 。

### 4. 启动项目

配置完成后，使用以下命令启动项目：

```bash
make run
```

### 5. 验证

打开浏览器或使用 curl 命令访问：

```
http://localhost:9000/api/ping
```

如果一切正常，应该看到 "pong" 响应。

## 项目结构

### `/cmd`

项目的主要应用程序。

- `/api`: API 服务器的入口点。

### `/config`

配置文件和配置加载逻辑。

- `config.go`: 主配置文件。
- `constant.go`: 常量定义。
- `environment_variable.go`: 环境变量处理。
- `router_white_list.go`: 路由白名单配置。

### `/doc`

设计文档、用户文档和其他项目相关文档。

### `/internal`

私有应用程序和库代码。

- `/api`: API 层，定义路由和处理函数。
- `/constant`: 内部使用的常量定义。
- `/database`: 数据库连接和初始化逻辑。
- `/middleware`: HTTP 中间件。
- `/model`: 数据模型和结构体定义。
- `/repository`: 数据访问层，处理数据持久化。
- `/service`: 业务逻辑层。

### `/migration`

数据库迁移文件。

- `000001_init_schema.down.sql`: 初始化 schema 的回滚脚本。
- `000001_init_schema.up.sql`: 初始化 schema 的执行脚本。

### `/pkg`

可以被外部应用程序使用的库代码。

- `/auth`: 认证相关功能。
- `/util`: 通用工具函数。

### `/test`

额外的外部测试应用程序和测试数据。

- `ping_test.go`: Ping 功能测试。
- `user_test.go`: 用户相关功能测试。

### 根目录文件

- `Makefile`: 定义常用命令的 Makefile。
- `project_bootstrap.sh`: 项目初始化脚本。

## 运行测试

本项目提供了多种运行测试的方式，以满足不同的测试需求。

### 运行所有测试

运行项目中的所有测试：

```bash
make test
```

### 运行特定文件夹中的测试

运行 ./test 文件夹中的测试：

```bash
make test-folder
```

### 运行特定测试文件

运行指定的测试文件：

```bash
make test-file FILE=path/to/your_test.go
```

例如：

```bash
make test-file FILE=./test/user_test.go
```

### 输出详细的测试日志

对于如上测试指令，都可加上`-verbose`参数，例如：

```bash
make test-verbose
make test-folder-verbose
make test-file-verbose FILE=./test/user_test.go 
```

## 切换运行环境

项目依赖于环境变量 `NODE_ENV` 来切换环境，不设置默认为 `development`。
项目会根据不同的环境加载不同的配置文件，如 `.env.development`、`.env.test`、`.env.production` 等。
只需创建一个新的 `.env` 文件，然后设置 `NODE_ENV` 变量即可。

## 更多使用方法

更多使用方法可以参考：

- `Makefile`文件，里面定义了常用的命令，如运行、测试、构建、清理等。
- 模块文件夹内的README文件
- 代码注释

后续将会持续更新和完善文档。

## 常见问题 (FAQ)

### Q1: 如何更改默认的服务器端口？

A: 在 `.env.development` 文件中修改 `SERVER_PORT` 变量。例如，将其设置为 `SERVER_PORT=8080` 会使服务器在 8080 端口上运行。

### Q2: 项目支持哪些数据库？

A: 目前项目主要支持 PostgreSQL 作为主数据库。如果您需要使用其他数据库，可能需要修改 `internal/database` 中的数据库连接代码。

### Q3: 如何添加新的 API 路由？

A: 在 `internal/api` 目录下相应的版本文件夹中（如 `v1`）添加新的处理函数，然后在 `internal/api/router.go` 文件中注册这个新路由。

### Q4: 如何自定义中间件？

A: 在 `internal/middleware` 目录下创建新的中间件文件，实现中间件逻辑，然后在 `internal/api/router.go` 中使用 `r.Use()`
方法应用这个中间件。

### Q5: 开发模式下的日志增强功能如何开启或关闭？

A: 这个功能默认在开发模式下自动开启。如果您想在生产环境中禁用它，请确保环境变量 `NODE_ENV` 不为 `development`。

### Q6: 如何运行特定的测试？

A: 使用 `make test-file` 命令，指定要运行的测试文件。例如：

```bash
make test-file FILE=./test/user_test.go
```

### Q7: JWT token 的过期时间如何配置？

A: 在 `pkg/auth/token.go` 文件中修改 `accessTokenExpire` 等变量。

### Q8: 如何添加新的自定义验证器？

A: 在 `pkg/util/formatValidator` 目录下添加新的验证函数，然后在 `pkg/util/customBindValidator` 中注册这个新的验证器。

### Q9: 项目是否支持 CORS？

A: 是的，项目默认配置了 CORS 中间件。

## 作者的话

前些年写Go项目的时候，自己总结了一个模板，最近开发新的项目，发现了一些不足，以及之前还有些不够完善的地方，一直没有系统地去更新和完善。
这次就直接从头开始，重新整理了一下，不仅便利了自己，还希望也能有机会帮到别的开发者。

这个项目会带有我的一些个人风格，比如目录结构、命名规范等，当然，这些都是可以根据自己的喜好和项目需求来调整的。如有建议，欢迎提Issue或PR。
