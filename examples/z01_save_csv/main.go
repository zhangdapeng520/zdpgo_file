package main

import "github.com/zhangdapeng520/zdpgo_file"

func main() {
	das := zdpgo_file.DoubleArrayString{
		Data: [][]string{
			{"a","b","c"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
			{"111","222","333"},
		},
	}
	das.SaveCsv("test.csv")
}