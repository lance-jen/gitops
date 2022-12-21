package math

import "testing"

func TestSomething(t *testing.T) {
    t.Skip()
}
func TestAdd(t *testing.T) {
    if Add(1, 2) == 3 {
        t.Log("math.Add PASS")
    } else {
        t.Error("math.Add FAIL")
    }
}
// func TestAdd2(t *testing.T) {
//     if Add(1, 2) == 2 {
//         t.Log("math.Add PASS")
//     } else {
//         t.Error("math.Add FAIL")
//     }
// }