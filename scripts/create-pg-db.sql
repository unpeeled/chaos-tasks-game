--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner:
--

CREATE TABLE tasks (
    id integer primary key generated always as identity,
    task text
);

--
-- Name: users; Type: TABLE; Schema: public; Owner: 
--

CREATE TABLE users (
    session uuid DEFAULT public.uuid_generate_v4(),
    name character varying(255),
    task_id integer
);


--
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: 
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- Name: users users_name_task_id_key; Type: CONSTRAINT; Schema: public; Owner: 
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_task_id_key UNIQUE (name, task_id);


--
-- Name: users users_session_key; Type: CONSTRAINT; Schema: public; Owner: 
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_session_key UNIQUE (session);


--
-- Name: users users_task_id_key; Type: CONSTRAINT; Schema: public; Owner:
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_task_id_key UNIQUE (task_id);


--
-- PostgreSQL database dump complete
--

