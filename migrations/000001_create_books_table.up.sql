CREATE TABLE IF NOT EXISTS books (
   id BIGSERIAL PRIMARY KEY,
   title TEXT NOT NULL,
   author TEXT NOT NULL,
   pages INTEGER NOT NULL,
   description TEXT,
   rating NUMERIC(3,2),
   genres TEXT[],
   created_at TIMESTAMP NOT NULL DEFAULT now(),
   updated_at TIMESTAMP NOT NULL DEFAULT now()
);
