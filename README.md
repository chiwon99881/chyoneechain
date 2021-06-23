## Chyoneechain with Go

- #01 Init

- #02 Singleton pattern

- #03 Create, Add, Get Blockchain

- #04 Set up for server

- #05 Templates for server html 1

- #06 Templates for server html 2

  > Use variable
  > Use partial, pages merge
  > get form data when post method

- #07 Refactoring for clean main.go

- #08 REST API with json

- #09 Interface

- #10 Handle GET / POST request

- #11 DefaultServeMux

  > ServeMux는 http request가 들어왔을 때, 이 아이가 어떤 handler를 실행시킬 지 파악을 한 후
  > 해당 handler를 연결 시켜주는 중간다리 역할을 하는 아이이다. 근데 이 아이가 커스텀으로 새로 지정하지 않고
  > ListenAndServe()에 대해 handler argument에 nil을 넣으면 DefaultServeMux를 사용하는데,
  > 이런 경우 포트가 다르더라도 같은 handler를 사용하기 때문에 같은 엔드포인트가 두 개 이상이면 에러를 발생시킨다.
  > 따라서 이번 commit은 그 에러를 해결하기 위해 나만의 커스텀 ServeMux를 만들었다.

- #12 Gorilla/mux

  > Gorilla/mux는 standard package가 아니라 third party package인데
  > ServeMux보다 효율이 좋다. 예를 들면, route의 변수를 사용해서 params를 가져오는 등
