package binarysearch

type Person struct {
	Name       string
	Surname    string
	Patronymic string
	Age        uint8
}

func Comparable(p1 *Person, p2 *Person) bool {
	return p1.Name < p2.Name && p1.Age < p2.Age
}
