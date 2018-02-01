package main

type Unicorn struct {
	Name string
}

func (u Unicorn) Rename(name string) {
	u.Name = name
}

func main() {
	larry := &Unicorn{
		Name: "Larry",
	}
	larry.Rename("sally")
	// What's larry.Name?
	// What can you change to modify that behavior?
}
