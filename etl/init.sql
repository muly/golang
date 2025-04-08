CREATE TABLE IF NOT EXISTS stage_student (
    id SERIAL PRIMARY KEY,
    name TEXT,
    address1 TEXT,
    address2 TEXT,
    city TEXT,
    state TEXT,
    zip TEXT,
    country TEXT
);

INSERT INTO stage_student (name, address1, address2, city, state, zip, country)
VALUES ('student1', 'address11', 'address21', 'city1', 'state', 'zip1', 'country1'),
       ('student2', 'address12', 'address22', 'city2', 'state', 'zip2', 'country2'),
       ('student3', 'address13', 'address23', 'city3', 'state', 'zip3', 'country3'),
       ('student4', 'address14', 'address24', 'city4', 'state', 'zip4', 'country4'),
       ('student5', 'address15', 'address25', 'city5', 'state', 'zip5', 'country5');


CREATE TABLE IF NOT EXISTS student (
    id SERIAL PRIMARY KEY,
    name TEXT,
    address TEXT
);

CREATE TABLE IF NOT EXISTS stage_school (
    id SERIAL PRIMARY KEY,
    name TEXT,
    address1 TEXT,
    address2 TEXT,
    city TEXT,
    state TEXT,
    zip TEXT,
    country TEXT
);

INSERT INTO stage_school (name, address1, address2, city, state, zip, country)
VALUES ('school1', 'address11', 'address21', 'city1', 'state', 'zip1', 'country1'),
       ('school2', 'address12', 'address22', 'city2', 'state', 'zip2', 'country2'),
       ('school3', 'address13', 'address23', 'city3', 'state', 'zip3', 'country3'),
       ('school4', 'address14', 'address24', 'city4', 'state', 'zip4', 'country4'),
       ('school5', 'address15', 'address25', 'city5', 'state', 'zip5', 'country5');

CREATE TABLE IF NOT EXISTS school (
    id SERIAL PRIMARY KEY,
    name TEXT,
    address TEXT
);