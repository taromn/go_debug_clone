package go_debug_clone

import "fmt"

/* 
 * [概要] 引数の型と値を出力する
 * [引数] 
 *   prefix: 前に挿入する接頭辞(string)
 *   printTarget:  型と値を出力する対象([T any])
 * [詳細]
 *   以下のような文字列がターミナルに出力される
 *   [<prefix>] 型: <printTargetの型>, 値: <printTargetの値>
 */
func PrintTypeAndValue[T any](prefix string, printTarget T) {
  fmt.Printf("[%s] 型: %T, 値: %+v\n", prefix, printTarget, printTarget)
}
