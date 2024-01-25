# Go-back

## Getting Started

To get "Go-back" up and running on your local development machine, follow these steps:

1. **Set Up the Live Reloader**:  
   Install the live-reloading tool to automatically compile your code when changes are detected:
   ```bash
   make setup-air
   ```

2. **Generate GraphQL Resolvers and Models**:  
   Set up the GraphQL layer by creating resolvers and models:
   ```bash
   make glg
   ```

3. **Database Migration**:  
   Apply the migrations to your database with Prisma:
   ```bash
   make migrate
   ```

4. **Start the Development Server**:  
   Boot up your application and start developing:
   ```bash
   make dev
   ```

Make sure your development environment is equipped with Go, and any other prerequisites mentioned in your project's documentation before running these commands.


## Description

"Go-back" is a robust backend framework built in Go, designed to streamline the development of web applications. It is built on a service-oriented architecture, ensuring a clean separation of concerns and modular development. The framework offers first-class support for GraphQL, allowing for the construction of flexible and efficient APIs that cater to clients' specific data requirements.

Key features include:

- **Prisma Integration**: Incorporates Prisma for type-safe and scalable database access, simplifying database operations and migrations with ease.
- **GraphQL Support**: Comes with a preconfigured GraphQL setup for developing flexible APIs, reducing the boilerplate needed to handle client requests.
- **Middleware Ready**: A ready-to-use middleware stack to handle authentication, error logging, request parsing, CORS, and other HTTP-related functionalities.
- **Router Layer**: Encapsulates business logic within routers, promoting clean code practices and making the codebase more maintainable.
- **Service Layer**: The service layer handle the logic behind database and third-party services. This separation of concerns ensures that business logic can be developed and maintained independently of the underlying data access or external service integrations.
- **Testing**: A dedicated testing suite to ensure the reliability and quality of the application through unit and integration tests.
