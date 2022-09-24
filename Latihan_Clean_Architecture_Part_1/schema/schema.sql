CREATE TABLE IF NOT EXISTS students (
    id varchar(5) not null primary key,
    nama varchar(100) not null,
    age int not null,
    grade int not null
);

INSERT INTO students (id,nama,age,grade) VALUES 
('B-001', 'Rifqi', 20, 100),
('B-002', 'Bruh', 21, 10),
('B-003', 'Udin', 22, 0)
;

