/*
pojo bean
*/
package pojo

type Cat struct {
	Id   int
	nick string
	age  int
}

func NewCat(id int) *Cat {
	return &Cat{Id: id}
}

func (c *Cat) SetNick(nick string) {
	c.nick = nick
}

func (c *Cat) SetAge(age int) {
	c.age = age
}

func (c *Cat) GetNick() string {
	return c.nick
}

func (c *Cat) GetAge() int {
	return c.age
}
