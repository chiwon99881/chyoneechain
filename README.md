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

- #13 Get a single block with gorillamux

- #14 Handling error with json response

- #15 Middleware and adapter pattern

- #16 CLI

  > os.Args는 command line에 입력한 인자를 받아오는 함수 예를 들어,
  > go run main.go AAA 라고 입력하면 [go run main.go, AAA]가 os.Args에 담기게 된다.

- #17 CLI Part 2 (flag)

- #18 CLI Part 3

  > Only used flag not FlagSet

- #19 Bolt for database

  > Bolt는 key:value쌍의 Go를 위한 Database이다.

- #20 Divide block & chain for Bolt database

- #21 Create and Save block in database

- #22 See what data saved with boltbrowser

  > go get github.com/br0xen/boltbrowser
  > go get github.com/evnix/boltdbweb

- #23 Restore our database Part (blockchain)

- #24 Restore our database Part (block) and find one block

- #25 Persist DONE

- #26 Mine blockchain (PoW)

- #27 Recalculate or Remain Difficulty

- #28 Recalculate or Remain Difficulty 2

- #29 PoW (Proof of Works) Done

- #30 Transaction Part 1

  > Reward to miner

- #31 Transaction Part 2

  > Balance after transactions

- #32 Transaction Part 3

  > Mempool and Unconfirmed transactions

- #33 Transaction Part 4

  > register transaction in mempool

- #34 Transaction Part 4

  > create block with transaction data

- #35 Transaction Part 5

  > Get unspent transaction outputs

- #36 Transaction Part 5

  > Transactions only using unspent transaction outputs

- #37 Transaction Part 6

  > Not allow coin if that coin is already in mempool

- #38 Transaction Part Done with refactoring

- #39 Digital Sign for Transaction

- #40 Restore privateKey

- #41 Singleton pattern for wallet

- #42 Singleton pattern for wallet Part 2

- #43 Sign, Verify signature Part 1

- #44 Sign, Verify signature Part 2

- #45 wallet Part DONE !

- #46 Upgrade http to WebSocket

- #47 P2P Connection Part 1

- #48 P2P Connection Part 2

- #49 P2P Connection Part 3

  > data races is occur when two or more goroutine access same data struct or data.
  > for example, when A goroutine edit \_DATA while when B goroutine delete \_DATA or read \_DATA.

- #50 P2P Connection Part 4 (Mutex)

  > modify error when occur data races by using Mutex.

- #51 P2P Connection Part 5 (send message)

- #52 Fixed data races

- #53 P2P Connection Part 6 (send message by message type)

- #54 P2P Connection Part 7 (replace all blocks by newest peer's blockchain)

- #55 P2P Connection Part 8 (broadcast new block)

- #56 P2P Connection Part 9 (broadcast new block 2)

- #57 P2P Connection Part 10

  > if one peer add new block, all the ohers peer's mempool also cleanup.
  > this commit do that.

- #58 P2P Connection Part 11 (Fix data races)

  > peers에 대해 읽고 있는 중에 peers에 inbox를 수정하려고 시도하기 때문에
  > data races가 일어나고 그 것을 수정

- #59 P2P Connection Part 12 (multi node part 1)

- #60 P2P Connection Part 13 (multi node part Done)

- #61 godocs

  > go get golang.org/x/tools/cmd/godoc
  > godoc -http:6060 -> http://localhost:6060 으로 가면
  > go docs같은 진짜 docs가 만들어진다 우리가 작업한 작업물에 대해서 !

- #62 Go Test

  > wannatestfilename_test.go

- #63 Go Test Part 2 (coverfile)

  > go test -v -coverprofile cover.out ./... (이거는 현재 test파일로 coverage가 얼마나 찍히는지까지 로그에 보여주고 cover.out파일을 생성한다.)

  > go tool cover -html=cover.out (이거는 cover.out파일을 html파일로 생성해서 보여줄 수 있는 명령어이다.)

- #64 Go Test Part 3 (how to test print log on console)

- #65 Go Test Part 4 (utils DONE)

  > go test -v -coverprofile cover.out ./... && go tool cover -html=cover.out
  > 이렇게 하면 html파일로 coverage를 보여주는 브라우저를 띄워주고, console에 pass,fail결과를 보여주는 테스트를 동시에 실행한다.

- #66 Go Test Part 5 (wallet Verify(), Sign())

- #67 Go Test Part 6 (prepare for test)

  > use interface and implement by create receiver function

- #68 Go Test Part 7 (TDD Test)

  > interface를 이용하여 test파일과 실제 go파일간에 원하는 결과를 처리해주기

- #69 Interface Power (adpater pattern)

  > interface를 통해서 구현하고자 하는 func을 정의해놓고 그 func의 signature를 그대로 사용하여 type또는 struct가 구현하면
  > test파일에서 임의의 struct를 같은 interface를 구현하여 테스트를 자유롭게 할 수 있음 왜그러냐면 테스트 파일에서는 실제 db에 접근하기를 원하는 건 아님
  > 그저 db에 접근을 통해서 가져오는 block, saveBlock, ... 등등을 가지고 우리가 구현한 코드를 검증하기위함인데
  > db에 접근하고 싶지 않기 때문에 fake struct를 같은 interface를 구현하여 가 데이터를 가져오고 지우고 하기만을 원함
  > 이게 첫번째 interface의 힘이고 두번째는 interface를 사용하면 추후에 다른 db를 사용한다한들 구현한 func외에 다른 파일을 수정할 필요가 없음
  > 그저 새로운 struct를 통해서 새로운 db에 대한 func만 구현해주면 됨 즉, 이와 같은 시그니쳐(dbStorage.findBlock)는 영원히 변하지 않을거란 얘기

- #70 Go Test Part 8 (CreateBlock func)
