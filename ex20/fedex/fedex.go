package fedex

import "fmt"

type FedexSender struct {
} 

// fedexSender가 가지고 있는 메서드 중에 Send라는 메서드
func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %s parcel\n", parcel)
} 