CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	telegram_id INTEGER NOT NULL UNIQUE,
	username TEXT NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);
