### Learning Goals
- Go - HTTP servers, structs, mutexes, JSON handling
- TypeScript - async/await, fetch API, DOM manipulation
- TailwindCSS

### Project Structure

```
STOP_WATCH/
├── backend/
│   ├── go.mod          # Go module file (like package.json for Node)
│   └── main.go         # Go server with extensive comments
├── frontend/
│   ├── src/
│   │   ├── main.ts     # TypeScript code with extensive comments
│   │   └── style.css   # TailwindCSS imports
│   └── index.html      # HTML with TailwindCSS classes explained
└── README.md
```
### Key Concepts

### Backend (Go)
- Structs: Grouping related data
- Mutex: Preventing race conditions
- HTTP Handlers: Responding to web requests
- JSON Encoding: Converting Go data to JSON
- CORS: Allowing frontend to access backend

### Frontend (TypeScript)
- async/await: Handling asynchronous operations
- fetch API: Making HTTP requests
- DOM manipulation: Changing HTTP elements
- setInterval: Repeating actions properly
- Event listenersL Responding to user clicks

### API endpoints 
- `GET /api/timer` - Get current timer status
- `Get /api/timer/start` - Start the timer
- `POST /api/timer/stop` - Stop the timer
- `POST /api/timer/reset` - Reset to 00:00.000
