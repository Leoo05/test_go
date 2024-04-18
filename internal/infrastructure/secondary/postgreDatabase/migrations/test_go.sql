CREATE TABLE course (
	id_course serial primary key,
	course_name varchar NOT NULL,	
);

CREATE TABLE student (
	id_student serial primary key,
	"name" varchar NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	document int NOT NULL,
	phone bigint NOT NULL,
	"address" varchar NOT NULL
);

CREATE TABLE grade (
	id_grade serial primary key,
	fk_student int NOT NULL,
	fk_course int NOT NULL,
	grade int NOT NULL,
	FOREIGN KEY(fk_student) REFERENCES student(id_student),
	FOREIGN KEY(fk_course) REFERENCES course(id_course)
);

CREATE TABLE professor (
	id serial primary key,
	"name" varchar NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	document int NOT NULL,
	fk_course int NOT NULL,
	FOREIGN KEY(fk_course) REFERENCES course(id_course)
);
