


#
make start-postgres
make clean-postgres start-postgres  


#
psql -h localhost -p 5432  -U postgres -d postgres

select * from stage_student;
select * from stage_school;

select * from student;
select * from school;

# 
go run main.go student.go school.go