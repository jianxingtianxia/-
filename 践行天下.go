package main

import "fmt"

//歌手的个人资料简介
type singer struct {
	name    string
	xingbie string
	Height  int
	weight  int
	age     int
}

func main() {
	caixvkun := singer{
		name:    "蔡徐坤",
		xingbie: "男",
		Height:  180,
		weight:  60,
		age:     24}
	fmt.Printf("我家哥哥是最棒的歌手\n")
	fmt.Printf("姓名:%s,性别:%s,身高:%dcm.体重:%dkg，年龄:%d岁\n", caixvkun.name, caixvkun.xingbie, caixvkun.Height, caixvkun.weight, caixvkun.age)

	fmt.Println(caixvkun)
	caixvkun.age += 1
	fmt.Printf("年龄：%d岁\n", caixvkun.age)

	wangjunkai := singer{
		name:    "王俊凯",
		xingbie: "男",
		Height:  182,
		weight:  58,
		age:     23}
	fmt.Printf("我家哥哥是最棒的歌手\n")
	fmt.Printf("姓名:%s,性别:%s,身高:%dcm,体重:%dkg,年龄:%d岁\n", wangjunkai.name, wangjunkai.xingbie, wangjunkai.Height, wangjunkai.weight, wangjunkai.age)
}
