# go-github-user-activity

A simple cli application that fetches a Github user's recent activity inspired by [Roadmap.sh](https://roadmap.sh/projects/github-user-activity).

## Building the application

```bash
go build -o build/github-user-activity *.go
```

## Running the application

### Development

```bash
go run *.go
```

### Production

```bash
// Build the application
go build -o build/github-user-activity *.go

// Run the application
./build/github-user-activity
```
