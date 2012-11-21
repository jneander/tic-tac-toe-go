package tictactoe

import "io"
import "strings"
import "strconv"

type ConsoleUI struct {
  in  Reader
  out Writer
}

func ( c ConsoleUI ) PromptMainMenu() {
  message := "Welcome to Tic Tac Toe in Go!\nPress any key to exit... "
  c.out.WriteString( message )
  for keys := ""; len(keys) == 0; keys = ReadLine( c.in ) {}
}

func ( c ConsoleUI ) DisplayBoard( b Board ) {
  spaces := b.Spaces()
  rows := make( []string, 3 )
  for i := range rows {
    rows[i] = strings.Join( spaces[i*3:i*3+3], "|" )
    rows[i] = strings.Replace( rows[i], b.Blank(), "_", -1 )
  }
  c.out.WriteString( strings.Join( rows, "\n" ) + "\n" )
}

func ( c ConsoleUI ) PromptPlayerMove( filter ...interface{} ) int {
  for {
    c.out.WriteString( "Please enter the space for your mark: " )
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
