-- Database: Survey

-- DROP DATABASE IF EXISTS "Survey";

CREATE DATABASE "Survey"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Turkish_Turkey.1254'
    LC_CTYPE = 'Turkish_Turkey.1254'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

BEGIN;


CREATE TABLE IF NOT EXISTS public.surveys
(
    token text COLLATE pg_catalog."default",
    title text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default"
);

CREATE TABLE IF NOT EXISTS public.options
(
    token text COLLATE pg_catalog."default",
    title text COLLATE pg_catalog."default",
    votes integer,
    surveytoken text COLLATE pg_catalog."default"
);

CREATE TABLE IF NOT EXISTS public.users
(
    token text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    perm text COLLATE pg_catalog."default" NOT NULL
);

CREATE TABLE IF NOT EXISTS public.votes
(
    token text COLLATE pg_catalog."default",
    optiontoken text COLLATE pg_catalog."default",
    usertoken text COLLATE pg_catalog."default",
    surveytoken text COLLATE pg_catalog."default"
);
END;
