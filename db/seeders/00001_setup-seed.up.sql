INSERT INTO dormitories(id, name) VALUES 
(1, 'Laura'), (2, 'Assystent');

INSERT INTO room_types(id, price, dormitory_id, name, max_of_residents) VALUES 
(1, 480, 1, 'two person room', 2), 
(2, 380, 2, 'two person room', 2),
(3, 305, 2, 'three person room', 3);

INSERT INTO rooms(id, room_number, room_type) VALUES 
(1, '415', 1),
(2, '19B', 2),
(3, '19A', 3);

INSERT INTO school_classes(id, name, class_year) VALUES
(1, 'TiB', 1),
(2, 'TiB', 2),
(3, 'TiB', 3),
(4, 'LogA', 1);

INSERT INTO supervisors(id, name, surname, email, phonenumber, dormitory_id) VALUES 
(1, 'viola', 'violovna', 'viola@gmail.com', '+123', 2),
(2, 'oscar', 'ocarov', 'oscar@gmil.com', '+124', 2),
(3, 'eva', 'evovna', 'eva@gamd.com', '+125', 1);

INSERT INTO pupils(id, name, surname, school_class_id, supervisor_id, room_id) VALUES 
(1, 'bogdana', 'boganova', 1, 1, 1),
(2, 'michael', 'lan', 3, NULL, 2),
(3, 'kalinin', 'kalin', 4, 3, 3);

INSERT INTO types_of_service (id, name) VALUES
(1, 'payment for lodging'),
(2, 'payment for superviser'), 
(3, 'payment for education');

INSERT INTO invoices(id, date_of_payment, pupil_id, type_of_service, amount_of_money, payment_due) VALUES 
(1, '2022-08-02', 1, 1, 0, '2022-09-02'),
(2, '2022-08-01', 2, 1, 380, '2022-09-02'),
(4, '2022-08-02', 3, 2, 500, '2023-08-02'),
(5, '2022-08-02', 2, 1, 380, '2022-10-02'), 
(6, '2022-04-02', 3, 1, 380, '2022-05-02'),
(7, '2022-04-02', 2, 1, 380, '2022-05-02');

INSERT INTO invoice_notes(invoice_id, note) VALUES 
(1, 'Bogdana did not pay because she have got help from our school for good grades');

/*
this is query to get everybody who owes Awans for lodging
SELECT invoices.pupil_id FROM invoices WHERE type_of_service=1 AND not CURRENT_DATE<payment_due AND invoices.pupil_id NOT IN (SELECT invoices.pupil_id FROM invoices WHERE CURRENT_DATE<payment_due AND type_of_service=1);
*/