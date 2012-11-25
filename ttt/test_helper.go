package ttt

func AddMarks( b *Board, mark string, set ...int ) {
  for _,p := range set {
    b.Mark( p, mark )
  }
}

type Queue struct {
  queue []interface{}
}

func (q *Queue) Push( vals ...interface{} ) {
  alt := make( []interface{}, len(q.queue) + len(vals) )
  copy( alt, q.queue )
  copy( alt[len(q.queue):], vals )
  q.queue = alt
}

func (q *Queue) Pop() (val interface{}) {
  if len(q.queue) > 0 {
    val, q.queue = q.queue[0], q.queue[1:]
    return val
  }
  panic( "Queue is empty and cannot be popped" )
}

func (q *Queue) IsEmpty() bool {
  return len(q.queue) == 0
}

func (q *Queue) Length() int {
  return len(q.queue)
}