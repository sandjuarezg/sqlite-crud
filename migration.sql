CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY,
    name TEXT,
    username TEXT,
    pass TEXT
);

INSERT INTO users (name, username, pass) 
    VALUES
        ('Marco Diaz','marco124', 'passMarco'),
        ('Dante Ramos','dante123', 'passDante');