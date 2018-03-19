package ttt

import (
	"io"
	"strconv"
	"strings"
)

type console struct {
	in  Reader
	out Writer
}

func NewConsole(in Reader, out Writer) *console {
	var ui = new(console)
	ui.in = in
	ui.out = out
	return ui
}

func (c console) PromptMainMenu() int {
	message := "\nWelcome to Tic Tac Toe in Go!\n" +
		"1) Player Goes First\n" +
		"2) Computer Goes First\n" +
		"3) Exit\n\n" +
		"Please enter your choice: "
	result := promptForInput(c, message, 1, 2, 3)
	switch result {
	case 1:
		return PLAYER_FIRST
	case 2:
		return COMPUTER_FIRST
	}
	return EXIT_GAME
}

func (c console) DisplayAvailableSpaces(b *Board) {
	rows := boardToASCII(b)
	vrows := availableSpacesToASCII(b)
	for i := range rows {
		rows[i] = "     " + rows[i] + "     " + vrows[i]
	}
	c.out.WriteString("\n" + strings.Join(rows, "\n") + "\n\n")
}

func (c console) DisplayBoard(b *Board) {
	rows := boardToASCII(b)
	for i := range rows {
		rows[i] = "     " + rows[i]
	}
	c.out.WriteString("\n" + strings.Join(rows, "\n") + "\n\n")
}

func (c console) DisplayGameOver(g Game) {
	winner, exists := g.Winner()
	if exists {
		c.out.WriteString("Player " + winner + " is the winner!\n\n")
	} else {
		c.out.WriteString("The game has ended in a draw!\n\n")
	}
}

func (c console) PromptPlayerMove(filter ...int) int {
	message := "Please enter the space for your mark: "
	for i := range filter {
		filter[i]++
	}
	return promptForInput(c, message, filter...) - 1
}

// PRIVATE

func promptForInput(c console, message string, filter ...int) int {
	for {
		c.out.WriteString(message)
		conv, err := strconv.Atoi(ReadLine(c.in))

		if err != nil {
			continue
		}

		if len(filter) == 0 {
			return conv
		} else if integerArrayPosition(filter, conv) > -1 {
			return conv
		}
	}
	return 0
}

func boardToASCII(board *Board) []string {
	rows := make([]string, 3)
	for i := range rows {
		rows[i] = strings.Join(board.Spaces()[i*3:i*3+3], "|")
		rows[i] = strings.Replace(rows[i], board.Blank(), "_", -1)
	}
	return rows
}

func availableSpacesToASCII(board *Board) []string {
	indices := make([]string, 9)
	for i := range indices {
		if board.Spaces()[i] == board.Blank() {
			indices[i] = strconv.Itoa(i + 1)
		} else {
			indices[i] = " "
		}
	}
	rows := make([]string, 3)
	for i := range rows {
		rows[i] = strings.Join(indices[i*3:i*3+3], " ")
	}
	return rows
}

func integerArrayPosition(array []int, element int) int {
	var pos = -1
	for i, v := range array {
		if v == element {
			pos = i
			break
		}
	}
	return pos
}

func ReadInput(buffer Reader) (result string) {
	var read = make([]byte, 128)
	num, _ := buffer.Read(read)
	return string(read[:num])
}

func ReadLine(reader Reader) string {
	var buffer = make([]byte, 1)
	var output string
	for {
		_, err := reader.Read(buffer)
		if buffer[0] == '\n' || err == io.EOF {
			break
		}
		output += string(buffer[0])
	}
	return output
}
