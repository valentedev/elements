CREATE TABLE IF NOT EXISTS elements_table
( 
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    atomic_number text COLLATE pg_catalog."default",
    symbol text COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    origin_of_name text COLLATE pg_catalog."default",
    periodo text COLLATE pg_catalog."default",
    grupo text COLLATE pg_catalog."default",
    block text COLLATE pg_catalog."default",
    atomic_weight text COLLATE pg_catalog."default",
    density_g_cm3 text COLLATE pg_catalog."default",
    melting_point text COLLATE pg_catalog."default",
    boiling_point text COLLATE pg_catalog."default",
    specific_heat_j_g text COLLATE pg_catalog."default",
    electro_negativity text COLLATE pg_catalog."default",
    abundance_in_earth_mg_kg text COLLATE pg_catalog."default",
    origin text COLLATE pg_catalog."default",
    phase_at_room_temperature text COLLATE pg_catalog."default"
)