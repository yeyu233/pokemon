package model

import (
	"fmt"
	"math/rand"
	"time"
)

//定义技能
type Skill struct {
	Name string
	Atk  int
}

//初始化技能
func NewSkill(name string, atk int) *Skill {
	return &Skill{
		Name: name,
		Atk:  atk,
	}
}

//定义精灵
type Pokemon struct {
	Name   string
	HP     int
	Skills [3]*Skill
}

//初始化精灵
func NewPokemon(name string, hp int, skill_list [3]*Skill) *Pokemon {
	return &Pokemon{
		Name:   name,
		HP:     hp,
		Skills: skill_list,
	}
}

//攻击动作，随机抽取一个技能，返回技能名称和攻击力，和被攻击后的精灵状态
func (baby *Pokemon) Attack(target Pokemon) (string, int, Pokemon) {
	skill := [3]*Skill{baby.Skills[0], baby.Skills[1], baby.Skills[2]}
	rand.Seed(int64(time.Now().UnixNano()))
	selected := skill[rand.Intn(3)]
	fmt.Println(baby.Name, "使用了 ", selected.Name)
	time.Sleep(1 * time.Second)
	target.HP = target.HP - selected.Atk
	fmt.Println(target.Name, "受到了", selected.Atk, "点伤害")
	fmt.Println(target.Name, "的生命值还有： ", target.HP)
	return selected.Name, selected.Atk, target
}
