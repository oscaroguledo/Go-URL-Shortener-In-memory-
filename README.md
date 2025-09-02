# Go-URL-Shortener-In-memory-
A URL shortener built with Go, using the Gin web framework and SQLite3 for persistent storage. Users can submit long URLs and receive short, unique slugs in return. The system supports redirection, link expiration, and in-memory caching for performance.

migrations:
# Create a new migration
./migrate.sh new create_urls_table

# Apply all migrations
./migrate.sh up

# Rollback one migration
./migrate.sh down 1

# Go to specific version
./migrate.sh goto 3


🔧 Key Features:

Accept long URLs and generate short slugs

Store URL mappings in a SQLite3 database

Serve redirects via GET /:slug

Optional expiration: automatically delete expired links

Simple REST API built with Gin

(Bonus) In-memory cache for fast lookups using a map[string]string