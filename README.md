# 2026 Countdown Application

A countdown application showing "Days left to complete 2026" with an orange, blocky Claude Code-inspired UI. Features a progress bar that fills daily and filters to view year/month/week progress.

## Tech Stack

- **Frontend**: React (JavaScript) + Vite
- **Backend**: Golang
- **Deployment**: Docker + Docker Compose

## Features

- Real-time countdown for 2026
- Filter by Year, Month, or Week
- Orange blocky theme (Claude Code style)
- Progress bar with percentage
- Auto-refresh every 60 seconds
- Fully responsive design

## Quick Start with Docker Compose

### Prerequisites

- Docker
- Docker Compose

### Run the Application

```bash
# Start the application
docker-compose up --build

# Or run in detached mode (background)
docker-compose up --build -d
```

### Access the Application

- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080

### Stop the Application

```bash
# Stop and remove containers
docker-compose down

# Stop, remove containers, and remove volumes
docker-compose down -v
```

## API Endpoints

- `GET /api/countdown` - Year progress (2026 completion)
- `GET /api/countdown/month` - Current month progress
- `GET /api/countdown/week` - Current week progress

## Manual Setup (Without Docker)

### Backend Setup

```bash
cd backend
go run main.go
```

Backend will start on http://localhost:8080

### Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

Frontend will start on http://localhost:5173

## Project Structure

```
countdown/
├── backend/               # Golang server
│   ├── handlers/         # API handlers
│   ├── utils/            # Date calculations
│   ├── main.go           # Server entry
│   ├── Dockerfile
│   └── go.mod
│
├── frontend/             # React app
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── hooks/       # Custom hooks
│   │   ├── services/    # API service
│   │   ├── App.jsx      # Root component
│   │   └── App.css      # Orange blocky theme
│   ├── Dockerfile
│   └── package.json
│
├── docker-compose.yml    # Docker Compose config
└── README.md
```

## Development

### Backend Development

```bash
cd backend
go run main.go
```

### Frontend Development

```bash
cd frontend
npm run dev
```

### Build Frontend for Production

```bash
cd frontend
npm run build
```

The backend will automatically serve the built frontend from the `dist/` folder.

## Design

**Orange Blocky Theme:**
- Primary color: #FF6B35
- No rounded corners (sharp, blocky)
- Thick borders (3-4px)
- Bold typography
- High contrast
- Box shadows for depth

## License

MIT
# GreatCountDown
# GreatCountDown
