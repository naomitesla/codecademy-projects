package main

import f "fmt"

func main(){
    var publisher, writer, artist, title string
    var year int16
    var pageNumber int8
    var grade float32

    title = "this course is too ez thus far :o"
    publisher = "naomi tesla c:"
    writer = "naomi :c"
    artist = "naomi :C"

    year = 2023
    pageNumber = 0
    grade = 100


    f.Println(title,
             "\n", "(", year, ",", pageNumber, ",", grade, ")", "\n",
             "published by", publisher, "\n",
             "also written by", writer, "\n",
             "yet again drawn by", artist, "\n")

    title = "hmmm, i wonder how i concat ints and strings without spaces, prob an inline type recast or something :o"
    pageNumber = 1
    grade = 50

    f.Println(title,
             "\n", "(", year, ",", pageNumber, ",", grade, ")", "\n",
             "published by", publisher, "\n",
             "also written by", writer, "\n",
             "yet again drawn by", artist)

}
