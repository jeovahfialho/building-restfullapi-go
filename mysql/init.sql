-- Ensure the 'bookdb' database exists, creating it if it does not
CREATE DATABASE IF NOT EXISTS bookdb;
USE bookdb;

-- Create the 'books' table if it does not already exist
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,       -- Unique identifier for each book
    title VARCHAR(255) NOT NULL,             -- Title of the book
    author VARCHAR(255),                     -- Author's name
    publishedYear INT,                       -- Year the book was published
    genre VARCHAR(100),                      -- Genre of the book
    summary TEXT,                            -- Summary or description of the book
    INDEX(title)                             -- Index on the 'title' column for faster searches
);

-- Grant full privileges on the 'bookdb' to the 'root' user from any host
GRANT ALL PRIVILEGES ON bookdb.* TO 'root'@'%' IDENTIFIED BY 'password';

-- Apply the privilege changes
FLUSH PRIVILEGES;
