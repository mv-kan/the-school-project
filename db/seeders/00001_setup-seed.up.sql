
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

INSERT INTO invoice_notes(id, note) VALUES 
(1, 'Bogdana did not pay because she have got help from our school for good grades');
/*
this is query to get everybody who owes Awans for lodging
SELECT invoices.pupil_id FROM invoices WHERE type_of_service=1 AND not CURRENT_DATE<payment_due AND invoices.pupil_id NOT IN (SELECT invoices.pupil_id FROM invoices WHERE CURRENT_DATE<payment_due AND type_of_service=1);
*/