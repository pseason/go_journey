package repo

import (
	"springmars.com/gin-mysql-viper/model"
)

// find by id
func FindStudentById(id int) (*model.Student, error) {
	var stu model.Student
	row := MysqlDb.QueryRow("select id, name, age from student where id=?", id)
	if err := row.Scan(&stu.Id, &stu.Name, &stu.Age); err != nil {
		return nil, err
	}
	return &stu, nil
}

// find all
func FindAllStudent(pageSize, pageNum int) ([]*model.Student, error) {
	rows, err := MysqlDb.Query("select * from student limit ?,?", (pageNum-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	// res
	stus := make([]*model.Student, 0)
	// 循环装配
	for rows.Next() {
		var stu model.Student
		if err := rows.Scan(&stu.Id, &stu.Name, &stu.Age); err != nil {
			return nil, err
		}
		stus = append(stus, &stu)
	}
	return stus, nil
}

// insert
func InsertStudent(stu model.Student) (int, error) {
	res, err := MysqlDb.Exec("insert into student(name,age)values(?,?)", stu.Name, stu.Age)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// update
func UpdateStudent(stu model.Student) (int, error) {
	res, err := MysqlDb.Exec("update student set name=?,age=? where id = ?", stu.Name, stu.Age, stu.Id)
	if err != nil {
		return 0, err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

// delete
func DeleteStudent(id int) (int, error) {
	res, err := MysqlDb.Exec("delete from student where id = ?", id)
	if err != nil {
		return 0, err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
