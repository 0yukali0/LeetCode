
type Stack struct {
    list []interface{}
}

func New() *Stack{
    s := new(Stack)
    s.list = make([]interface{}, 0, 8)
    return s
}

func (s *Stack) IsEmpty() bool{
    return len(s.list) == 0
}

func (s *Stack) Pop() interface{}{
    if s.IsEmpty(){
        return nil
    }
    tmp := s.list[len(s.list) - 1]
    s.list = s.list[:len(s.list)-1]
    return tmp
}

func (s *Stack) Push(element interface{}){
    s.list = append(s.list, element)
}

func isValid(s string) bool {
   stack := New()
   pairs := map[rune]rune{
       '}': '{',
       ']': '[',
       ')': '(',
   }
   for _, c := range s {
       if c == '}' || c == ']' || c == ')' {
           if stack.IsEmpty() {
               return false;
           }
           got := stack.Pop().(rune)
           expect := pairs[c]
           if expect != got {
               return false
           }
       } else {
           stack.Push(c)
       }
   }

   if !stack.IsEmpty() {
       return false
   }
   
   return true
}
