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
    room_type_id INT NOT NULL,
    UNIQUE(room_number, room_type_id),
    FOREIGN KEY (room_type_id) REFERENCES room_types(id)
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
    type_of_service_id INT NOT NULL,
    amount_of_money DECIMAL NOT NULL,
    payment_start DATE NOT NULL,
    payment_due DATE NOT NULL,
    FOREIGN KEY (pupil_id) REFERENCES pupils(id),
    FOREIGN KEY (type_of_service_id) REFERENCES types_of_service(id)
);

CREATE TABLE IF NOT EXISTS invoice_notes (
    invoice_id INT PRIMARY KEY REFERENCES invoices(id),
    note TEXT NOT NULL
); 

INSERT INTO dormitories(id, name) VALUES 
(DEFAULT, 'Laura'), (DEFAULT, 'Assystent');

INSERT INTO room_types(id, price, dormitory_id, name, max_of_residents) VALUES 
(DEFAULT, 480, 1, 'two person room', 2), 
(DEFAULT, 380, 2, 'two person room', 2),
(DEFAULT, 305, 2, 'three person room', 3);

INSERT INTO rooms(id, room_number, room_type_id) VALUES 
(DEFAULT, '415', 1),
(DEFAULT, '19B', 2),
(DEFAULT, '19A', 3);

INSERT INTO school_classes(id, name, class_year) VALUES
(DEFAULT, 'TiB', 1),
(DEFAULT, 'TiB', 2),
(DEFAULT, 'TiB', 3),
(DEFAULT, 'LogA', 1);

INSERT INTO supervisors(id, name, surname, email, phonenumber, dormitory_id) VALUES 
(DEFAULT, 'viola', 'violovna', 'viola@gmail.com', '+123', 2),
(DEFAULT, 'oscar', 'ocarov', 'oscar@gmil.com', '+124', 2),
(DEFAULT, 'eva', 'evovna', 'eva@gamd.com', '+125', 1);

INSERT INTO pupils(id, name, surname, school_class_id, supervisor_id, room_id) VALUES 
(DEFAULT, 'bogdana', 'boganova', 1, 1, 1),
(DEFAULT, 'michael', 'lan', 3, NULL, 2),
(DEFAULT, 'kalinin', 'kalin', 4, 3, 3);

INSERT INTO types_of_service (id, name) VALUES
(DEFAULT, 'payment for lodging'),
(DEFAULT, 'payment for superviser'), 
(DEFAULT, 'payment for education');

INSERT INTO invoices(id, date_of_payment, pupil_id, type_of_service_id, amount_of_money, payment_start, payment_due) VALUES 
(DEFAULT, '2022-08-02', 1, 1, 0,'2022-08-02', '2022-09-02'),
(DEFAULT, '2022-08-01', 2, 1, 380,'2022-08-01', '2022-09-02'),
(DEFAULT, '2022-08-02', 3, 2, 500,'2022-08-02', '2023-08-02'),
(DEFAULT, '2022-08-02', 2, 1, 380,'2022-08-02', '2022-10-02'), 
(DEFAULT, '2022-04-02', 3, 1, 380,'2022-04-02', '2022-05-02'),
(DEFAULT, '2022-04-02', 2, 1, 380,'2022-04-02', '2022-05-02');

INSERT INTO invoice_notes(invoice_id, note) VALUES 
(1, 'Bogdana did not pay because she have got help from our school for good grades');