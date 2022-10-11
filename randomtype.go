package randomname

type RandomType int

const (
	AdjectiveAndPerson RandomType = iota // 形容词+人物 例如：帅气的罗纳尔多
	PersonActSomething                   // 人物+做事情 例如：梅西吃爆米花
)
