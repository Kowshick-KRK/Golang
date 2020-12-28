package main
 
import (
   "encoding/json"
   "fmt"
   "log"
   "Playground/avl"
)
 
type intKey int
 
func (k intKey) Less(k2 avl.Key) bool { return k < k2.(intKey) }
func (k intKey) Eq(k2 avl.Key) bool   { return k == k2.(intKey) }
 
func dump(tree *avl.Node) {
   b, err := json.MarshalIndent(tree, "", "   ")
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(string(b))
}
 
func main() {
   var tree *avl.Node
   fmt.Println("Empty tree:")
   dump(tree)
 
   fmt.Println("\nInsert test:")
   avl.Insert(&tree, intKey(4))
   avl.Insert(&tree, intKey(5))
   avl.Insert(&tree, intKey(1))
   avl.Insert(&tree, intKey(2))
   dump(tree)
 
   fmt.Println("\nRemove test:")
   avl.Remove(&tree, intKey(4))
   avl.Remove(&tree, intKey(1))
   dump(tree)
}
 