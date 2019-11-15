package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"pokemon/model"
	"strconv"
	"time"
)

var (
	Input      string
	Str        string
	Skill01    model.Skill
	Skill02    model.Skill
	Skill03    model.Skill
	skill_list [3]*model.Skill
	Player01   model.Pokemon
	Player02   model.Pokemon
	Player     model.Pokemon
)

func NewPlayer(name string, skill01 string, skill02 string, skill03 string) *model.Pokemon {
	Skill01 := model.NewSkill(skill01, 2)
	Skill02 := model.NewSkill(skill02, 3)
	Skill03 := model.NewSkill(skill03, 5)
	skill_list = [3]*model.Skill{Skill01, Skill02, Skill03}
	Player := model.NewPokemon(name, 30, skill_list)
	return Player

}

func ShowPlayerInfo(Player01 model.Pokemon) {
	fmt.Println("您的精灵是：")
	fmt.Println(Player01.Name, "HP: ", Player01.HP, "技能: ", Player01.Skills[0].Name, Player01.Skills[1].Name, Player01.Skills[2].Name)
}
func ShowAiInfo(Player01 model.Pokemon) {
	fmt.Println("您的对手是：")
	fmt.Println(Player01.Name, "HP: ", Player01.HP, "技能: ", Player01.Skills[0].Name, Player01.Skills[1].Name, Player01.Skills[2].Name)
}

func choicePokemon(Str string) model.Pokemon {
	var Player model.Pokemon
	if Str == "1" {
		Player := NewPlayer("皮卡丘", "摇尾巴", "电光一闪", "十万伏特")

		return *Player
	} else if Str == "2" {
		Player := NewPlayer("胖丁", "唱歌", "拍击", "连环巴掌")

		return *Player
	} else if Str == "3" {
		Player := NewPlayer("可达鸭", "念力", "石化功", "疯狂瘙痒")

		return *Player
	} else if Str == "4" {
		Player := NewPlayer("超梦", "幻象术", "波导弹", "精神破坏")

		return *Player
	} else if Str == "5" {
		Player := NewPlayer("波克比", "飞撞", "高压水枪", "自我暗示")

		return *Player
	} else {
		fmt.Println("选个已有的精灵好不")
		os.Exit(0)
	}
	return Player
}

func randomC() int {
	rand.Seed(int64(time.Now().UnixNano()))
	r := rand.Intn(5)
	if r == 0 {
		r = rand.Intn(5)
		return r
	}
	return r
}

func main() {

	fmt.Println("选择你的精灵：")
	fmt.Println("1 皮卡丘，2 胖丁，3 可达鸭，4 超梦，5 波克比")
	f := bufio.NewReader(os.Stdin)
	Input, _ = f.ReadString('\n')
	fmt.Sscan(Input, &Str)
	Player01 := choicePokemon(Str)
	fmt.Println(" =======Start======= ")
	ShowPlayerInfo(Player01)
	re := randomC()
	for strconv.Itoa(re) == Input {
		re = randomC()
	}
	Player02 := choicePokemon(strconv.Itoa(re))
	ShowAiInfo(Player02)
	time.Sleep(2 * time.Second)

	round := 1
	for 1 > 0 {

		fmt.Println("======= Round", round, " =======")
		fmt.Println("你的回合：")
		_, _, Player02 = Player01.Attack(Player02)
		if Player02.HP <= 0 {
			fmt.Println(Player01.Name, "获得了胜利！")
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("对手回合：")
		_, _, Player01 = Player02.Attack(Player01)
		if Player01.HP <= 0 {
			fmt.Println(Player02.Name, "获得了胜利！")
			break
		}

		time.Sleep(2 * time.Second)
		round = round + 1
	}

}
