Blockchain tutorial:
===================

[Source Link](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)  
[Source Code](https://github.com/Jeiwan/blockchain_go)

Usage:
------

Remove my.db if it exists:

```
go run *.go getbalance -address Ivan
go run *.go send -from Ivan -to Pedro -amount 6
go run *.go getbalance -address Ivan
go run *.go getbalance -address Pedro
go run *.go send -from Pedro -to Helen -amount 2
go run *.go send -from Ivan -to Helen -amount 2
go run *.go send -from Helen -to Rachel -amount 3
go run *.go getbalance -address Ivan
go run *.go getbalance -address Pedro
go run *.go getbalance -address Helen
go run *.go getbalance -address Rachel
go run *.go send -from Pedro -to Ivan -amount 5
```
