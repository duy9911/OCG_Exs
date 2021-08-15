

# Crawler program


## Description 

This exercice which implements a program crawl films from https://www.imdb.com/chart/top/?ref_=nv_mv_250 

### Main packages
packge grom
packge colly 

## Process 

1 connect db.
2 create function to crawl with  a channel as a sender  
3 create function save these value in DB with a channel as a receiver
4 improve performance with goroutine 
5 await two goroutine before end!

## Usage
```go
cd src path 
go run main.go
```

## References 
Connect db: 
https://gorm.io/docs/connecting_to_the_database.html
Goroutine:
https://golangbot.com/goroutines/
https://www.google.com/search?q=waitgroup&sxsrf=ALeKk01uf1ryloE8ZT6TsdpXR9uW27wLiA%3A1629038303237&ei=3yYZYZPvDdTBhwOvpoZA&oq=waitgroup&gs_lcp=Cgdnd3Mtd2l6EAMyBQgAEIAEMgUIABCABDIFCAAQywEyBQgAEMsBMgUIABDLATIFCAAQywEyBQgAEMsBMgUIABDLATIFCAAQywEyBQgAEMsBOgcIABBHELADOgQIIxAnOgsIABCABBCxAxCDAToICAAQgAQQsQM6BAgAEAM6CAguELEDEIMBOgYIIxAnEBM6DgguEIAEELEDEMcBEKMCOgUILhCABDoICC4QgAQQsQM6BwgAEAoQywFKBAhBGABQjwtY2xZggxhoAnACeACAAYABiAHZCJIBAzEuOZgBAKABAcgBCMABAQ&sclient=gws-wiz&ved=0ahUKEwjT2o7LoLPyAhXU4GEKHS-TAQgQ4dUDCA4&uact=5





