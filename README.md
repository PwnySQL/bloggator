# bloggator
Simple CLI tool to aggregate updates of RSS feeds, e.g. from blogs, in a SQL database

## TODO
1. Add pagination to browse command
2. Add more filters (e.g. feeds) to browse command
3. Remove HTML tags from title, description, ... using [bluemonday](https://github.com/microcosm-cc/bluemonday?tab=readme-ov-file)
4. Add a service for the agg command to run it in the background and restarts it if it crashes
5. Add search command with fuzzy searching through posts
6. Add a TUI to read a post in the terminal

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `bloggator` with:

```bash
go install ...
```

## Config

Create a `.bloggatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
bloggator register <name>
```

Add a feed:

```bash
bloggator addfeed <url>
```

Start the aggregator:

```bash
bloggator agg 30s
```

View the posts:

```bash
bloggator browse [limit]
```

There are a few other commands you'll need as well:

- `bloggator login <name>` - Log in as a user that already exists
- `bloggator users` - List all users
- `bloggator feeds` - List all feeds
- `bloggator follow <url>` - Follow a feed that already exists in the database
- `bloggator unfollow <url>` - Unfollow a feed that already exists in the database

