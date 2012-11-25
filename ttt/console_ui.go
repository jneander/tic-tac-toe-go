package ttt

import "io"
import "strings"
import "strconv"

type ConsoleUI struct {
  in  Reader
  out Writer
}

func NewConsoleUI( in Reader, out Writer ) *ConsoleUI {
  var ui = new( ConsoleUI )
  ui.in = in
  ui.out = out
  return ui
}

func ( c ConsoleUI ) PromptMainMenu() int {
  message := "Welcome to Tic Tac Toe in Go!\n" +
             "Enter one of the following options:\n" +
             "1) Player vs Player\n" +
             "2) Exit\n\n"
  result := promptForInput( c, message, 1, 2 )
  return result
}

func ( c ConsoleUI ) DisplayAvailableSpaces( b *Board ) {
  rows := boardToASCII( b )
  vrows := availableSpacesToASCII( b )
  for i := range rows {
    rows[i] = "     " + rows[i] + "     " + vrows[i]
  }
  c.out.WriteString( "\n" + strings.Join( rows, "\n" ) + "\n\n" )
}

func ( c ConsoleUI ) DisplayBoard( b *Board ) {
  rows := boardToASCII( b )
  for i := range rows { rows[i] = "     " + rows[i] }
  c.out.WriteString( "\n" + strings.Join( rows, "\n" ) + "\n\n" )
}

func ( c ConsoleUI ) PromptPlayerMove( filter ...interface{} ) int {
  message := "Please enter the space for your mark: "
  return promptForInput( c, message, filter... ) - 1
}

func promptForInput( c ConsoleUI, message string, filter ...interface{} ) int {
  for {
    c.out.WriteString( message )
    conv,err := strconv.Atoi( ReadLine( c.in ) )

    if err != nil { continue }

    if len( filter ) == 0 {
      return conv
    } else if arrayPosition( filter, conv ) > -1 {
      return conv
    }
  }
  return 0
}

func boardToASCII( board *Board ) []string {
  rows := make( []string, 3 )
  for i := range rows {
    rows[i] = strings.Join( board.Spaces()[i*3:i*3+3], "|" )
    rows[i] = strings.Replace( rows[i], board.Blank(), "_", -1 )
  }
  return rows
}

func availableSpacesToASCII( board *Board ) []string {
  indices := make( []string, 9 )
  for i := range indices {
    if board.Spaces()[i] == board.Blank() {
      indices[i] = strconv.Itoa(i + 1)
    } else {
      indices[i] = " "
    }
  }
  rows := make( []string, 3 )
  for i := range rows {
    rows[i] = strings.Join( indices[i*3:i*3+3], " " )
  }
  return rows
}

func arrayPosition( array []interface{}, element interface{} ) int {
  var pos = -1
  for i,v := range array {
    if v == element {
      pos = i
      break
    }
  }
  return pos
}

func ReadInput( buffer Reader ) ( result string ) {
  var read = make( []byte, 128 )
  num, _ := buffer.Read( read )
  return string( read[:num] )
}

func ReadLine( reader Reader ) string {
  var buffer = make( []byte, 1 )
  var output string
  for {
    _,err := reader.Read( buffer )
    if buffer[0] == '\n' || err == io.EOF { break }
    output += string( buffer[0] )
  }
  return output
}
