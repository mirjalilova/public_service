CREATE TABLE Party (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name VARCHAR NOT NULL,
                       slogan VARCHAR,
                       opened_date DATE,
                       description TEXT,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW(),
                       deleted_at BIGINT DEFAULT 0
);

CREATE TABLE Public (
                        id VARCHAR PRIMARY KEY,
                        first_name VARCHAR NOT NULL,
                        last_name VARCHAR NOT NULL,
                        birthday DATE,
                        gender VARCHAR,
                        nation VARCHAR,
                        party_id VARCHAR,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW(),
                        deleted_at BIGINT DEFAULT 0,
                        FOREIGN KEY (party_id) REFERENCES Party(id)
);


