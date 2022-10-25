# desafio-cierre-testing

Proyecto Base para cumplir el Desafio de Testing

![image](https://user-images.githubusercontent.com/114087997/197588709-a5e35f67-fa6b-4ed1-90b2-6267751eeb82.png)


## Changelog

Mon Oct 24 10:50:27 -03 2022
- remove "go.mod" and "go.sum"
- `go mod init example.com` (the project is inside the _"backpack"_)
- fix imports accordingly
- new "go.mod" and "go.sum" with `go mod tidy`
- update README :blush:

Mon Oct 24 11:44:38 -03 2022
- API tested with __Insomnia: OK__
<img width="943" alt="insomniaValidation" src="https://user-images.githubusercontent.com/114087997/197555382-5d748a73-fbdd-48fe-acd6-cee33f7de0c6.png">

Mon Oct 24 12:43:31 -03 2022 (Test Service)
- add "service_test.go"
- add "testify"
- no mocks nor stubs needed: Repository is already mocked
- neither Service nor Repository check for valid SellerID, no need to check for `err != nil`
- query validation is only done in the handler
- test: `err == nil`: :thumbsup:
- test: `data == mockedData`: :thumbsup:

Mon Oct 24 13:40:20 -03 2022
- add "mockedRepository" to test error response
- test for errors and empty data
- coverage so far is __50%__ 

Mon Oct 24 13:42:40 -03 2022
- add "handler_test.go"
- create fake server
- create request and response recorder
- coverage has reached __80%__ :thumbsup:

Mon Oct 24 14:22:11 -03 2022
- create failServer using NewMockedRepository to test read fail
- ¡coverage has reached __100%__ :blush: !
