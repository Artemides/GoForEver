package store

import "fmt"

type Fruits struct {
	id   string
	name string
	cost float64
}

func (F *Fruits) SetName(name string) {
	F.name = name
}
func (F *Fruits) SetId(id string) {
	F.id = id
}
func (F *Fruits) SetCost(cost float64) {
	F.cost = cost
}
func (F *Fruits) Id() string {
	return F.id
}
func (F *Fruits) Name() string {
	return F.name
}

func (F *Fruits) Cost() float64 {
	return F.cost
}
func (F *Fruits) SetValues(id, name string, cost float64) {
	F.id = id
	F.name = name
	F.cost = cost
}
func (F Fruits) String() string {
	return fmt.Sprintf("%s,%s,%.2f", F.Id(), F.Name(), F.Cost())
}
