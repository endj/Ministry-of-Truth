CREATE TABLE IF NOT EXISTS user_profiles (
    id SERIAL PRIMARY KEY,
    external_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    info VARCHAR(240) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_templates (
    id SERIAL PRIMARY KEY,
    user_profile_id INT NOT NULL,
    name TEXT NOT NULL,
    age INT,
    occupation TEXT,
    personality JSONB,
    interests JSONB,
    background JSONB,
    communication_style JSONB,
    social_connections JSONB,
    CONSTRAINT fk_user_profiles
        FOREIGN KEY (user_profile_id)
        REFERENCES user_profiles (id)
        ON DELETE CASCADE
);
