package main

import "fmt"

func main() {
	items := map[string]map[string][]string{}

	// Add some values
	items["Ivan"] = map[string][]string{
		"Books":     []string{"Cisco", "Patterns", "Algorithms"},
		"Magazines": []string{"Play Boy"},
	}

	items["Semen"] = map[string][]string{
		"Books":     []string{"Chemistry"},
		"Magazines": []string{"Young naturalist"},
	}

	// Show number of users
	fmt.Println("Number of users:", len(items))

	// Show how many books and magazines user has
	users := make([]string, 0, len(items))
	for u := range items {
		fmt.Print(u, ": ")
		users = append(users, u)
		count := len(items[u]["Books"]) + len(items[u]["Magazines"])
		fmt.Println(count, "books and magazines")
	}
}
