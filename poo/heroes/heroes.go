package heroes

type Hero struct {
	id    int
	name  string
	speed float64
	life  float64
	build []string
}

func NewHero(id int, name string, speed, life float64, build []string) *Hero {
	return &Hero{
		id:    id,
		name:  name,
		speed: speed,
		life:  life,
		build: build,
	}
}
