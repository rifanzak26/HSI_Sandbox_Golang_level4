-- This file contains the initial SQL schema for the database.

-- department table.
CREATE TABLE departments (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) UNIQUE
);

-- employee table with foreign key.
CREATE TABLE employee (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100),
  department_id INT NOT NULL,
  -- foreign key constraint
  FOREIGN KEY (department_id) REFERENCES departments(id)
);