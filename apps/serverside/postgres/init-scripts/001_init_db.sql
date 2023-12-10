CREATE TABLE public.items (
    id serial PRIMARY KEY,
    name text NOT NULL,
    quantity integer,
    description text
);
