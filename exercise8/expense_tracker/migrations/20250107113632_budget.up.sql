CREATE TABLE IF NOT EXISTS budget (
    id_budget SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW() ,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (id_user) REFERENCES users_(id_user)
);