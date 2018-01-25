package handle

import (
  "AT_Message_Server/handle/ios"
)

// 想象一下，如果一个系统中有一万个接口，并且时刻在工作，那会是怎样的一个系统
func Manager() {
  ios.Handles()
}

