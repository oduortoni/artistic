# Rtstk

A web application built with Go (Beego) backend and TypeScript + Vite frontend.

## Description

Rtstk is a website for artistic people to showcase their creative works. More like the github for artists, only there are no commandline tools to "push" your art, but who knows, maybe we could one day get artists to push pictures on display. That was a joke by the way.

## Technologies in use

Rtstk (read artistic) is a modern web application that integrates a Go backend using the Beego framework with a TypeScript frontend powered by Vite. The application features hot-reloading for both frontend and backend development, making it ideal for rapid development cycles.

## Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- npm 9 or higher

## Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/oduortoni/rtstk.git
   cd rtstk
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Install the Air tool for Go hot-reloading:
   ```bash
   go install github.com/air-verse/air@latest
   ```

4. Install frontend dependencies:
   ```bash
   cd frontend
   npm install
   ```

5. Start the development servers:
   ```bash
   npm run dev
   ```

This will start:
- Frontend development server on `http://localhost:5173`
- Backend API server on `http://localhost:8000`
- Hot-reloading for both frontend and backend changes

## Production Deployment

1. Build the frontend:
   ```bash
   cd frontend
   npm run build
   ```
   This will create optimized production files in `frontend/dist`

2. Build the Go application:
   ```bash
   go build -o rtstk
   ```

3. Run the production server:
   ```bash
   ./rtstk
   ```

The production server will serve:
- The static frontend files from the `frontend/dist` directory
- The API endpoints on port 8000

## Project Structure

```
rtstk/
├── controllers/     # Backend API controllers
├── models/         # Database models
├── routers/        # API route definitions
├── backend/        # Backend with the domain driven concerns
├── frontend/       # Frontend application
│   ├── src/        # TypeScript source files
│   └── dist/       # Production build output
└── main.go         # Application entry point
```

## Contributing

1. Fork the repository
2. Create your feature branch:
   ```bash
   git checkout -b feature/my-new-feature
   ```
3. Commit your changes:
   ```bash
   git commit -am 'Add some feature'
   ```
4. Push to the branch:
   ```bash
   git push origin feature/my-new-feature
   ```
5. Submit a pull request

### Development Guidelines

- Follow Go best practices and formatting guidelines
- Write tests for new features
- Update documentation as needed
- Ensure both frontend and backend changes are tested together

## License

MIT
 