---
applyTo: '**'
---

# CentralAuth Instructions
CentralAuth is a Single Sign-On (SSO) solution for my personal projects. It allows users to log in once and access multiple applications without needing to log in again for each application.

# Tech Stack
## Client
- React with TypeScript
- Tanstack Router for file-based routing
- Tanstack Query for data fetching
- Zustand for state management
- Shadcn/UI for component library
- Axios for API requests
- Vite as build tool

## Server
- Go backend
- echo for HTTP server
- SQLC for type-safe SQL
- Goose for database migrations
- JWT for authentication

# Project Structure
The project follows a domain-driven design with clear separation of concerns:

## Client Structure
- `/api`: API integration layer with domain-specific subfolders
    - Each domain has `.api.ts` (API calls), `.dao.ts` (data types), and `.query.ts` (Tanstack Query)
- `/components`: UI components organized by page/feature
- `/hooks`: Custom React hooks
- `/lib`: Utility functions and configurations
- `/routes`: Tanstack Router file-based routing
- `/stores`: Zustand state management stores
- `/utils`: Helper functions

## Server Structure
- `/cmd`: Application entry points
- `/internal`: Internal packages
    - `/config`: Application configuration
    - `/db`: Database connection and management
        - `/migrations`: Goose SQL migrations
        - `/queries`: SQLC query definitions
        - `/sqlc`: Generated SQLC code
    - `/features`: Feature-specific code
        - `/auth`: Authentication logic
            - `auth.handler.go`: HTTP handler for authentication
            - `registeration.handler.go`: registeration entrypoint
            - `login.handler.go`: login entrypoint
    - `/middlewares`: HTTP middleware functions
    - `/utils`: Helper functions

# Coding Conventions
- Follow the given design principles (currently using a simplified version without service and moldel layers because it is done by SQLC)
- Maintain clear separation of concerns
- Use TypeScript interfaces for type safety on the client
- Follow Go idioms and conventions on the server
- Organize code by features rather than by type
- Use Tanstack Query for all API data fetching
- Implement proper error handling at all levels
- Just Implement those feature that have been discussed and agreed upon


# Goal
The goal of this project is to create a simple and efficient SSO solution that can be easily integrated into various applications. The project should be modular, maintainable, and scalable to accommodate future growth and additional features.( it won't be using PKCE anymore, but it will be using a custom JWT implementation)

for now focus only on the backend api side only 