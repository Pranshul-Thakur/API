# Album API

A simple REST API built with Go and Gin framework for managing album records.

## Features

- GET all albums
- GET album by ID
- POST new album

## Installation

```bash
go mod download
```

## Usage

Run the server:
```bash
go run main.go
```

The API will be available at `http://localhost:8081`

## Endpoints

**GET /albums** - Retrieve all albums

**POST /albums** - Add a new album
```json
{
  "id": "4",
  "title": "Album Title",
  "artist": "Artist Name",
  "price": 29.99
}
```

**GET /albums/:id** - Get album by ID
