package database

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func NewDB(databaseURL string) (*DB, error) {
	log.Printf("DEBUG: Original database URL: %s", databaseURL)
	
	// Resolve hostname to IPv4 to avoid IPv6 issues
	databaseURL = resolveToIPv4(databaseURL)
	
	log.Printf("DEBUG: Resolved database URL: %s", databaseURL)
	
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Printf("ERROR: Failed to open database connection: %v", err)
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	log.Printf("DEBUG: Attempting to ping database...")
	if err := conn.Ping(); err != nil {
		log.Printf("ERROR: Database ping failed: %v", err)
		log.Printf("DEBUG: Connection string used (password visible): %s", databaseURL)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	log.Printf("DEBUG: Database connection successful!")

	db := &DB{conn: conn}
	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}

func (d *DB) GetConn() *sql.DB {
	return d.conn
}

func (d *DB) Close() error {
	return d.conn.Close()
}

func resolveToIPv4(databaseURL string) string {
	// Extract hostname from postgresql:// URL
	// Format: postgresql://user:pass@host:port/db?params
	if !strings.HasPrefix(databaseURL, "postgresql://") {
		return databaseURL
	}
	
	parts := strings.Split(databaseURL, "@")
	if len(parts) != 2 {
		return databaseURL
	}
	
	hostPart := parts[1]
	hostAndPort := strings.Split(hostPart, "/")[0]
	hostPortParts := strings.Split(hostAndPort, ":")
	
	if len(hostPortParts) < 2 {
		return databaseURL
	}
	
	hostname := hostPortParts[0]
	port := hostPortParts[1]
	
	// Resolve to IPv4
	ips, err := net.LookupIP(hostname)
	if err != nil {
		log.Printf("Warning: Failed to resolve %s, using original URL: %v", hostname, err)
		return databaseURL
	}
	
	var ipv4 net.IP
	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4 = ip
			break
		}
	}
	
	if ipv4 == nil {
		log.Printf("Warning: No IPv4 address found for %s, using original URL", hostname)
		return databaseURL
	}
	
	// Replace hostname with IPv4
	newURL := strings.Replace(databaseURL, hostname+":"+port, ipv4.String()+":"+port, 1)
	log.Printf("Resolved %s to IPv4: %s", hostname, ipv4.String())
	
	return newURL
}

func (d *DB) migrate() error {
	schemas := []string{
		bootcampSchema,
		journalSchema,
		memeSchema,
		storySchema,
		contactSchema,
		chatSchema,
	}

	for _, schema := range schemas {
		if _, err := d.conn.Exec(schema); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	log.Println("Database migrations completed successfully")
	return nil
}

const bootcampSchema = `
CREATE TABLE IF NOT EXISTS bootcamps (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    subtitle TEXT NOT NULL,
    description TEXT NOT NULL,
    long_description TEXT NOT NULL,
    duration TEXT NOT NULL,
    level TEXT NOT NULL,
    price TEXT NOT NULL,
    tech_stack TEXT NOT NULL,
    highlights TEXT NOT NULL,
    project_features TEXT NOT NULL,
    target_audience TEXT NOT NULL,
    images TEXT,
    videos TEXT,
    github_repo TEXT,
    demo_url TEXT,
    status TEXT NOT NULL CHECK(status IN ('active', 'upcoming', 'completed')),
    enrolled_count INTEGER DEFAULT 0,
    rating REAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bootcamp_modules (
    id TEXT PRIMARY KEY,
    bootcamp_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    duration TEXT NOT NULL,
    topics TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bootcamp_id) REFERENCES bootcamps(id) ON DELETE CASCADE
);
`

const journalSchema = `
CREATE TABLE IF NOT EXISTS journal_entries (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    summary TEXT NOT NULL,
    published_on DATE NOT NULL,
    category TEXT,
    tags TEXT,
    author TEXT,
    read_time TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const memeSchema = `
CREATE TABLE IF NOT EXISTS meme_categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS memes (
    id TEXT PRIMARY KEY,
    category_id TEXT NOT NULL,
    type TEXT NOT NULL CHECK(type IN ('img', 'yt', 'mp4', 'webm')),
    src TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES meme_categories(id) ON DELETE CASCADE
);
`

const storySchema = `
CREATE TABLE IF NOT EXISTS stories (
    id TEXT PRIMARY KEY,
    media TEXT NOT NULL,
    mimetype TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const contactSchema = `
CREATE TABLE IF NOT EXISTS contacts (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    message TEXT NOT NULL,
    subject TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const chatSchema = `
CREATE TABLE IF NOT EXISTS chat_sessions (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chat_messages (
    id TEXT PRIMARY KEY,
    session_id TEXT NOT NULL,
    role TEXT NOT NULL CHECK(role IN ('user', 'assistant', 'system')),
    content TEXT NOT NULL,
    reasoning_details TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (session_id) REFERENCES chat_sessions(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_chat_messages_session_id ON chat_messages(session_id);
CREATE INDEX IF NOT EXISTS idx_chat_messages_created_at ON chat_messages(created_at DESC);
`

