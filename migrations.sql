CREATE TABLE IF NOT EXISTS books (
    id serial primary key,
    title text not null,
    description text not null,
    author text not null
);
