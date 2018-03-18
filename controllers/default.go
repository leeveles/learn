package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    // var a [3]int
    // fmt.Println(a[0])
    // fmt.Println(a[len(a)-1])

    // for i,v := range a {
    //     fmt.Printf("%d %d",i,v)
    // }

    // for _,v := range a {
    //     fmt.Printf("%d", v)
    // }
/*
    var a [3]int = [3]int{0,1,2}

    fmt.Printf("%d", a[1])

    var b [3]int = [3]int{0,1}

    fmt.Printf("%d", b[2])
*/
    //type Currency int

    // const (
    //     USD Currency = iota
    //     EUR 
    //     GBP
    //     RMB
    // )

    // symbol := [...]string{USD:"$", EUR:"E", GBP:"XX", RMB:"7"}

    // fmt.Println(RMB, symbol[RMB])

    fmt.Printf("\n\n\n\n\n\n");

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
    c.TplName = "learn.tpl"
}
