package main

import (
	"fmt"
	"github.com/gongwalker/parseIni"
)

func main()  {
	c,err := parseIni.ReadIniFile("./demo/demo.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 默认获取配置值为string型 (此函数的另一个别名函数为 GetConfigToString )
	strvalue,err := c.GetConfig("mysql.passwd") 
	fmt.Printf("value=%v \t type=%T \t err=%v \n",strvalue,strvalue,err)
	// 输出 value=abc123 	 type=string 	 err=<nil>


	// 获取配置值且转为int型
	value,err := c.GetConfigToInt("mysql.port")
	fmt.Printf("value=%v \t type=%T \t err=%v \n",value,value,err)
	// 输出 value=3306 	 type=int 	 err=<nil>

	// 获取不存在的key值
	value,err = c.GetConfigToInt("noneKey")
	fmt.Printf("value=%v \t type=%T \t err=%v \n",value,value,err)
	// 输出 value=0 	 type=int 	 err=指定的key不存在

	// 获取配置值且转为float型
	fvalue,err := c.GetConfigToFloat64("pi")
	fmt.Printf("value=%v \t type=%T \t err=%v \n",fvalue,fvalue,err)
	// 输出 value=3.1415 	 type=float64 	 err=<nil>

	// 获取配置值且转为bool型 (非"0"便为true)
	bvalue,err := c.GetConfigToBool("isCache")
	fmt.Printf("value=%v \t type=%T \t err=%v \n",bvalue,bvalue,err)
	// 输出 value=true 	 type=bool 	 err=<nil>

	// 获取配置值，会忽略配置文件两边的 " 与 '
	svalue,err := c.GetConfigToString("省会.辽宁")
	fmt.Printf("value=%s \t type=%T \t err=%v \n",svalue,svalue,err)
	// 输出 value=沈阳 	 type=string 	 err=<nil>

	svalue,err = c.GetConfigToString("省会.陕西")
	fmt.Printf("value=%s \t type=%T \t err=%v \n",svalue,svalue,err)

	svalue,err = c.GetConfigToString("省会.河北")
	fmt.Printf("value=%s \t type=%T \t err=%v \n",svalue,svalue,err)
	// 输出 value=Shi Jia Zhuang 	 type=string 	 err=<nil>
}


