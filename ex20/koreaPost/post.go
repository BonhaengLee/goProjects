package koreaPost

import "fmt"
// 우체국 택배가 제공하는 api(패키지)
type PostSender struct {
} 
// 가지고 있는 메서드
func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %s를 보냅니다.\n", parcel)
} 