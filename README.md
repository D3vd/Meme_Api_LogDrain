# Meme API LogDrain

Log Processor for Meme_Api
Acts as an HTTPs LogDrain using Heroku's LogPlex and stores API Analytics to MongoDB on a day by day basis

## Running the App

### 1. Setup the App

Clone the repo and run the following commands to install all necessary dependencies and build the app

```bash
  go install
  go build .
```

### 2. Set Environment Variables

Set the following Environment Variables

```bash
  export MONGODB_URI=<MONGODB_URI_HERE>
  export SENTRY_DSN=<SENTRY_DSN_HERE>
```

### 3. Run the App

Run the app by executing the binary

```bash
  ./Meme_Api_LogDrain
```

### 4. Setup HTTPs LogDrain for the Heroku App

Setup LogDrain for your Heroku App using the Heroku CLI

```bash
  heroku drains:add https://<APP_HOSTED_URL>/log -a myapp
```
