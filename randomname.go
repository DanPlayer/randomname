package randomname

import (
	"math/rand"
	"time"
)

// GenerateName 生成随机昵称
func GenerateName() string {
	var name string
	rand.Seed(time.Now().UnixNano())
	selectedType := RandomType(rand.Intn(1))
	switch selectedType {
	case AdjectiveAndPerson:
		name = AdjectiveSlice[rand.Intn(AdjectiveSliceCount)] + PersonSlice[rand.Intn(PersonSliceCount)]
	case PersonActSomething:
		name = PersonSlice[rand.Intn(PersonSliceCount)] + ActSomethingSlice[rand.Intn(ActSomethingSliceCount)]
	}
	return name
}
