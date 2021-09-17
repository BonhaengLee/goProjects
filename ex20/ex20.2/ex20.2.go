package main

import "goProjects/ex20/koreaPost"

// 만약 우체국 택배로 바꾸면 코드 수정 비용 발생(산탄총 코드)
// 따라서 어떤 회사를 쓰든 Sender를 통해 쓰겠다. >> 인터페이스를 사용하면 된다.
type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {//sender *fedex.FedexSender) { 
	sender.Send(name);
}

func main() {
	// var sender Sender = &fedex.FedexSender{} // 해당 인스턴스를 sender에 할당
	// 인스턴스 사용시 만약 회사가 바뀐다면 호출하는 인스턴스와 메서드만 바꿔주면 된다.
	 var sender Sender = &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
