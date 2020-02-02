- Declaration of new variable: use `var` or `:=`
`var card string = "Ace of Spades"`
`card := "Ace of Spades"`

- Reassigning value: use `=`
`card = "Ace of Spades"`

- print
`fmt.Println(card)`

- Declaring a function
`func newCard() string { } `

- Data Structure
  1. Array: Fixed length list of things
  2. Slice: An array that can grow or shrink

- Creating a new Slice of string
` cards := []string{"Ace of Diamonds", newCard()} `

- Slice indices



- for loop
```
for index,card := range cards {
  fmt.Println(card)
}
```



