package model

type Customer struct {
	Name string
	Age  int
}

var Customers []Customer

func (c *Customer) Add(datas ...map[string]interface{}) {
	for _, v := range datas {
		c.Name = v["name"].(string)
		c.Age = v["age"].(int)
	}

	Customers = append(Customers, *c)
}

func (c *Customer) Edit() {

}

func (c *Customer) Delete() {

}
