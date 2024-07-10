CREATE TABLE if NOT EXISTS product(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(8,2),
    available BOOLEAN, 
    created TIMESTAMP DEFAULT NOW()    
)