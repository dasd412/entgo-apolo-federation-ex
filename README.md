1️⃣ User Service
* 사용자 정보 관리 (회원가입, 로그인, 프로필)
* JWT 기반 인증
***
* 2️⃣ Order Service
* 주문(Order) 생성, 조회
* 주문 상태 변경 (Pending, Paid, Shipped, Canceled)
***
* 3️⃣ Delivery Service
* 배송(Delivery) 요청 처리
* 배송 상태 변경 (Pending, In Transit, Delivered)
  서비스 간 연동 흐름
***
1. 회원 가입 (User Service)
    * User 생성 (id 반환)
2. 주문 생성 (Order Service)
    * Order 생성 (user_id, total_price, status=Pending 저장)
3. 배송 요청 (Delivery Service)
    * Delivery 생성 (order_id, status=Pending)
4. 배송 진행 (Delivery Service)
    * 배송 상태가 In Transit, Delivered로 변경
    * Order 상태도 Shipped로 변경
