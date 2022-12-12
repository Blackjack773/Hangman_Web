package main

import "fmt"

// Tableau de string representant chacun une lettre en Ascii-Art

func Lettertoascii(s string) []string {
	tab := []string{}
	switch s {
	case " ":
		str := []string{"  ", "  ", "  ", "  ", "  ", "  ", "  "}
		return str
	case "a":
		str := []string{"           ", "    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "          "}
		return str
	case "b":
		str := []string{" ____   ", "|  _ \\  ", "| |_) | ", "|  _ <  ", "| |_) | ", "|____/  ", "          "}
		return str
	case "c":
		str := []string{"  _____  ", " / ____| ", "| |      ", "| |      ", "| |____  ", " \\_____| ", "          "}
		return str
	case "d":
		str := []string{" _____   ", "|  __ \\  ", "| |  | | ", "| |  | | ", "| |__| | ", "|_____/  ", "          "}
		return str
	case "e":
		str := []string{" ______  ", "|  ____| ", "| |__    ", "|  __|   ", "| |____  ", "|______| ", "          "}
		return str
	case "f":
		str := []string{" ______  ", "|  ____| ", "| |__    ", "|  __|   ", "| |      ", "|_|      ", "          "}
		return str
	case "g":
		str := []string{"  _____  ", " / ____| ", "| |  __  ", "| | |_ | ", "| |__| | ", " \\_____| ", "          "}
		return str
	case "h":
		str := []string{" _    _  ", "| |  | | ", "| |__| | ", "|  __  | ", "| |  | | ", "|_|  |_| ", "          "}
		return str
	case "i":
		str := []string{" _____  ", "|_   _| ", "  | |   ", "  | |   ", " _| |_  ", "|_____| ", "          "}
		return str
	case "j":
		str := []string{"      _  ", "     | | ", "     | | ", " _   | | ", "| |__| | ", " \\____/  ", "          "}
		return str
	case "k":
		str := []string{" _  __ ", "| |/ / ", "| ' /  ", "|  <   ", "| . \\  ", "|_|\\_\\ ", "          "}
		return str
	case "l":
		str := []string{" _       ", "| |      ", "| |      ", "| |      ", "| |____  ", "|______| ", "          "}
		return str
	case "m":
		str := []string{" __  __  ", "|  \\/  | ", "| \\  / | ", "| |\\/| | ", "| |  | | ", "|_|  |_| ", "          "}
		return str
	case "n":
		str := []string{" _   _  ", "| \\ | | ", "|  \\| | ", "| . ` | ", "| |\\  | ", "|_| \\_| ", "          "}
		return str
	case "o":
		str := []string{"  ____   ", " / __ \\  ", "| |  | | ", "| |  | | ", "| |__| | ", " \\____/  ", "          "}
		return str
	case "p":
		str := []string{" _____   ", "|  __ \\  ", "| |__) | ", "|  ___/  ", "| |      ", "|_|      ", "          "}
		return str
	case "q":
		str := []string{"  ____   ", " / __ \\  ", "| |  | | ", "| |  | | ", "| |__| | ", " \\___\\_\\ ", "          "}
		return str
	case "r":
		str := []string{" _____   ", "|  __ \\  ", "| |__) | ", "|  _  /  ", "| | \\ \\  ", "|_|  \\_\\ ", "          "}
		return str
	case "s":
		str := []string{"  _____  ", " / ____| ", "| (___   ", " \\___ \\  ", " ____) | ", "|_____/  ", "          "}
		return str
	case "t":
		str := []string{" _______  ", "|__   __| ", "   | |    ", "   | |    ", "   | |    ", "   |_|    ", "          "}
		return str
	case "u":
		str := []string{" _    _  ", "| |  | | ", "| |  | | ", "| |  | | ", "| |__| | ", " \\____/  ", "          "}
		return str
	case "v":
		str := []string{"__      __ ", "\\ \\    / / ", " \\ \\  / /  ", "  \\ \\/ /   ", "   \\  /    ", "    \\/     ", "          "}
		return str
	case "w":
		str := []string{"__          __ ", "\\ \\        / / ", " \\ \\  /\\  / /  ", "  \\ \\/  \\/ /   ", "   \\  /\\  /    ", "    \\/  \\/     ", "          "}
		return str
	case "x":
		str := []string{"__   __ ", "\\ \\ / / ", " \\ V /  ", "  > <   ", " / . \\  ", "/_/ \\_\\ ", "          "}
		return str
	case "y":
		str := []string{"__     __ ", "\\ \\   / / ", " \\ \\_/ /  ", "  \\   /   ", "   | |    ", "   |_|    ", "          "}
		return str
	case "z":
		str := []string{" ______ ", "|___  / ", "   / /  ", "  / /   ", " / /__  ", "/_____| ", "          "}
		return str
	default:
		fmt.Printf("\n")
		fmt.Println(Red + "Veuillez entrer un caractère correcte svp" + Reset)
		fmt.Printf("\n")
	}
	return tab
}
