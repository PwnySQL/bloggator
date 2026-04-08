# bloggator
Simple CLI tool to aggregate updates of RSS feeds, e.g. from blogs, in a SQL database

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

