DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'working_model_enum') THEN
        CREATE TYPE working_model_enum AS ENUM ('in-person', 'hybrid', 'remote');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    role TEXT NOT NULL,
    description TEXT NOT NULL,
    requirements_qualifications JSON NOT NULL,
    additional_information TEXT,
    working_model working_model_enum NOT NULL,
    process_steps JSON NOT NULL,
    company TEXT NOT NULL,
    location TEXT NOT NULL,
    salary INTEGER NOT NULL,
    link TEXT NOT NULL,
    publish_date DATE DEFAULT CURRENT_DATE,
    expiration_date DATE NOT NULL
);
