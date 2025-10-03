# Deploy IoT API to Vercel

This guide shows how to deploy your Go IoT API to Vercel for global access.

## 🚀 Quick Deploy

### 1. Install Vercel CLI
```bash
npm i -g vercel
```

### 2. Login to Vercel
```bash
vercel login
```

### 3. Deploy
```bash
vercel --prod
```

## 📁 Project Structure

```
iot-device/
├── api/
│   ├── health.go      # Health check endpoint
│   └── data.go        # Data receiving endpoint
├── vercel.json        # Vercel configuration
├── main.go           # Original local server (keep for local dev)
├── go.mod
└── README.md
```

## 🌐 Available Endpoints

After deployment, your API will be available at `https://your-project.vercel.app`:

- `GET /health` - Health check
- `POST /data` - Receive IoT data from Raspberry Pi

## 🔧 Environment Variables

If you need database functionality, set these in Vercel dashboard:

1. Go to [Vercel Dashboard](https://vercel.com/dashboard)
2. Select your project
3. Go to Settings → Environment Variables
4. Add:
   - `DB_HOST`
   - `DB_PORT` 
   - `DB_USER`
   - `DB_PASSWORD`
   - `DB_NAME`

## 📱 Update Raspberry Pi Code

Update your Python code to use the Vercel URL:

```python
# Change this line in rpb.py
server_url = "https://your-project.vercel.app/data"
```

## 🧪 Test Your Deployment

```bash
# Test health endpoint
curl https://your-project.vercel.app/health

# Test data endpoint
curl -X POST https://your-project.vercel.app/data \
  -H "Content-Type: application/json" \
  -d '{"device_id": "test", "temperature": 25.5}'
```

## 🔄 Local Development

For local development, continue using:
```bash
go run main.go
```

This runs the full server with database connectivity on `localhost:8080`.

## 📝 Notes

- Vercel functions are serverless - they spin up on demand
- Each API endpoint is a separate function
- Database connections should use connection pooling for serverless
- Logs are available in the Vercel dashboard

## 🚨 Important

Make sure to update your Raspberry Pi's `server_url` to point to your new Vercel deployment URL!