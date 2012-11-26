package ttt

type Player interface {
  Move( Board ) int
  SetMark( string )
  GetMark() string
}
