-- Create the uuid-ossp extension in the public schema
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
-- Grant usage on the extension functions to public
GRANT USAGE ON SCHEMA public TO PUBLIC;