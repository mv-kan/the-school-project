/*
create domain email with regular expression for email validation
*/
CREATE EXTENSION citext;
CREATE DOMAIN email AS citext
  CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );

CREATE TABLE IF NOT EXISTS dormitories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(500) UNIQUE NOT NULL 
);

CREATE TABLE IF NOT EXISTS room_types (
    id SERIAL PRIMARY KEY,
    price MONEY NOT NULL,
    dormitory_id INT NOT NULL,
    name VARCHAR(500) NOT NULL,
    max_of_residents INT NOT NULL,
    FOREIGN KEY (dormitory_id) REFERENCES dormitories(id)
);

CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY,
    room_number VARCHAR(50) NOT NULL,
    room_type INT NOT NULL,
    UNIQUE(room_number, room_type),
    FOREIGN KEY (room_type) REFERENCES room_types(id)
);

CREATE TABLE IF NOT EXISTS school_classes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL, 
    class_year INT NOT NULL,
    UNIQUE(name, class_year)
);

CREATE TABLE IF NOT EXISTS supervisors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    surname VARCHAR(500) NOT NULL,
    email email UNIQUE NOT NULL,
    phonenumber VARCHAR(50) UNIQUE NOT NULL, 
    dormitory_id INT NOT NULL,
    FOREIGN KEY (dormitory_id) REFERENCES dormitories(id)
);

CREATE TABLE IF NOT EXISTS pupils (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(500) NOT NULL,
    surname VARCHAR(500) NOT NULL,
    email email, 
    phonenumber VARCHAR(50), 
    school_class_id INT NOT NULL, 
    /*
    Pupils may or may not have supervisor or room provided by Awans
    */
    supervisor_id INT, 
    room_id INT,
    FOREIGN KEY (school_class_id) REFERENCES school_classes(id),
    FOREIGN KEY (supervisor_id) REFERENCES supervisors(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)   
);

CREATE TABLE IF NOT EXISTS types_of_service (
    id SERIAL PRIMARY KEY,
    name VARCHAR(500) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS invoices(
    id SERIAL PRIMARY KEY,
    date_of_payment DATE NOT NULL,
    pupil_id INT NOT NULL,
    type_of_service INT NOT NULL,
    amount_of_money MONEY NOT NULL,
    payment_due DATE NOT NULL,
    FOREIGN KEY (pupil_id) REFERENCES pupils(id),
    FOREIGN KEY (type_of_service) REFERENCES types_of_service(id)
);

CREATE TABLE IF NOT EXISTS invoice_notes (
    invoice_id INT PRIMARY KEY REFERENCES invoices(id),
    note TEXT NOT NULL
); 

