package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id      int
	name    string
	country string
	created time.Time
}

func main() {
	var employee Employee
	employee.id = 1
	employee.name = "John Doe"
	employee.country = "USA"
	employee.created = time.Now()
	fmt.Printf("%+v\n", employee)
}
