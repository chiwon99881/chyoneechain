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
