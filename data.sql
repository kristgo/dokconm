CREATE TABLE subscriber (

    id int(11) not null,
    firstname varchar(20) not null,
    lastname varchar(20) not null,
    birthyear int not null,
    PRIMARY KEY(id)
);

INSERT INTO subscriber values 
(1,'James', 'Hadley', '1972'),
(2,'Susan', 'Tingley', '1969'),
(3,'Frank', 'Tucker', '1988'),
(4,'John', 'Hayward', '1976'),
(5,'Charles', 'Frasier', '1991'),
(6,'John', 'Spencer', '2003'),
(7,'Richard', 'Farley', '1968'),
(8,'Stanley', 'Cheswick', '1999'),
(9,'Carl', 'Lutz', '2002'),
(10,'Paul', 'Hardy', '1984'),
(11,'John', 'Chilton', '1995'),
(12,'Thomas', 'Ziegler', '1976');