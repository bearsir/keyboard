// get user input number

package keyboard

import (
  "bufio"
  "os"
  "strconv"
  "strings"
)

// 取得輸入
func GetFloat() (float64, error) {
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if err != nil {
    return 0, err
  }
  // 轉換大小寫
  input = strings.TrimSpace(input)
  number, err := strconv.ParseFloat(input, 64)
  if err != nil {
    return 0, err
  }
  return number, err
}

