-- Database: ecommerce
CREATE DATABASE IF NOT EXISTS ecommerce;
USE ecommerce;

-- Table: Users
CREATE TABLE Users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	email VARCHAR(100) NOT NULL UNIQUE,
	password_hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Table: Categories
CREATE TABLE Categories (
	category_id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: Products
CREATE TABLE Products (
	product_id INT AUTO_INCREMENT PRIMARY KEY,
	category_id INT,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	price DECIMAL(10, 2) NOT NULL,
	stock INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	FOREIGN KEY (category_id) REFERENCES Categories(category_id) ON DELETE SET NULL
);

-- Table: Orders
CREATE TABLE Orders (
	order_id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	order_date DATE NOT NULL,
	total_amount DECIMAL(10, 2) NOT NULL,
	status ENUM('Pending', 'Shipped', 'Delivered', 'Cancelled') DEFAULT 'Pending',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);


-- Table: Addresses (One-to-Many relationship with Users)
CREATE TABLE Addresses (
	address_id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	street VARCHAR(255) NOT NULL,
	city VARCHAR(100) NOT NULL,
	state VARCHAR(50),
	postal_code VARCHAR(20),
	country VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);


-- Table: Tags
CREATE TABLE Tags (
	tag_id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL UNIQUE
);

-- Table: ProductTags (Many-to-Many relationship between Products and Tags)
CREATE TABLE ProductTags (
	product_id INT NOT NULL,
	tag_id INT NOT NULL,
	PRIMARY KEY (product_id, tag_id),
	FOREIGN KEY (product_id) REFERENCES Products(product_id) ON DELETE CASCADE,
	FOREIGN KEY (tag_id) REFERENCES Tags(tag_id) ON DELETE CASCADE
);

-- Table: Reviews (One-to-Many relationship with Products and Users)
CREATE TABLE Reviews (
	review_id INT AUTO_INCREMENT PRIMARY KEY,
	product_id INT NOT NULL,
	user_id INT NOT NULL,
	rating INT CHECK (rating BETWEEN 1 AND 5),
	comment TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (product_id) REFERENCES Products(product_id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
