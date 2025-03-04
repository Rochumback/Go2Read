CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username  varchar(100) UNIQUE NOT NULL,
    email varchar(100) UNIQUE NOT NULL,
    password_hash char(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS genre (
    id serial PRIMARY KEY,
    genre_name varchar(30) NOT NULL UNIQUE
);

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE OR REPLACE FUNCTION generate_searchable(title varchar, alternative_titles varchar[])
    RETURNS TEXT AS $$
    BEGIN
    RETURN title || array_to_string(alternative_titles, '', '*');
END;
$$ LANGUAGE plpgsql IMMUTABLE;

CREATE TABLE IF NOT EXISTS manga (
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    alternative_titles varchar(255)[],
    searchable text GENERATED ALWAYS AS ( generate_searchable(title, alternative_titles) ) STORED,
    synopsis text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    created_by int NOT NULL DEFAULT 0 REFERENCES users ( id ) ON DELETE SET DEFAULT 
);

CREATE TABLE IF NOT EXISTS manga_genres (
    manga int REFERENCES manga(id) ON DELETE CASCADE,
    genre int REFERENCES genre(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS chapter (
    id serial PRIMARY KEY,
    views_count int DEFAULT 0,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    pages_count smallint NOT NULL,
    chapter_number decimal NOT NULL,
    manga_id int NOT NULL REFERENCES manga(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS index_searchable_mangas ON public.manga USING gist (searchable public.gist_trgm_ops (siglen='64') );
CREATE INDEX IF NOT EXISTS index_chapters_manga ON public.chapter USING hash ( manga_id );
