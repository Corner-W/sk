package main

import (
	"fmt"
	"github.com/Corner-W/sk/util/sqnum"

	"os"
	"time"
)

func main() {
	n, err := sqnum.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node:", id.Node(),
			"step:", id.Step(),
			"time:", id.Time(),
			"\n",
		)
		tm := time.Unix(id.Time(), 0)

		fmt.Println(tm.Format("2006-01-02 15:04:05"))
	}

}
