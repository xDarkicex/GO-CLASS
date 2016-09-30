package main

import "fmt"

func main() {
	nested()
}

func nested() {

	for {
		for i := 0; i < 2; i++ {
			fmt.Printf("%s%d\n", "Loop 1: ", i)
			for i := 0; i < 3; i++ {
				fmt.Printf("%s%d\n", "Loop 2: ", i)
				for i := 0; i < 4; i++ {
					fmt.Printf("%s%d\n", "Loop 3: ", i)
					for i := 0; i < 5; i++ {
						fmt.Printf("%s%d\n", "Loop 4: ", i)
						for i := 0; i < 6; i++ {
							fmt.Printf("%s%d\n", "Loop 5: ", i)
							for i := 0; i < 7; i++ {
								fmt.Printf("%s%d\n", "Loop 6: ", i)
								for i := 0; i < 8; i++ {
									fmt.Printf("%s%d\n", "Loop 7: ", i)
									for i := 0; i < 9; i++ {
										fmt.Printf("%s%d\n", "Loop 8: ", i)
										for i := 0; i < 10; i++ {
											fmt.Printf("%s%d\n", "Loop 9: ", i)
											for i := 0; i < 11; i++ {
												fmt.Printf("%s%d\n", "Loop 10: ", i)
											}
										}
									}
								}
							}
						}
					}
					break
				}
				break
			}
			break
		}
		fmt.Println("I am done")
		break
	}
}
