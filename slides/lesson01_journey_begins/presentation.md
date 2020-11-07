# The way of the Gopher <br><small>the journey begins</small>

------

# Why?

## Code is the pen and paper of today; literacy is important
## It is art, and can be very rewarding
## You've tried Python, and it sucks
## Money? Median [salary is $140k]

### Cute animals?

![Gopher](https://blog.golang.org/gopher/header.jpg)

(the Go Gopher! credit to Renee French, [borrowed from here])

[salary is $140k]:  https://www.businessinsider.com/highest-paying-programming-languages-stack-overflow-developer-survey-2020-5#2-go-8
[borrowed from here]: https://blog.golang.org/gopher

---

# This series

## Teaches you Go, along with other important tools of the 2020's

## At the end, you can scrape the web or something

## We will use martial arts as an analogy

1. Learn to pick up the blade
2. Learn to swing the blade
3. Learn the battleground
4. Finally, command your forces

## Beyond the videos, you will need to practice every day. Hone your skills.

---

# This lesson

## Pick up your equipment (download our development environment)

1. Download [VirtualBox] ([windows] and [OS X]; Linux users know what to do)
2. Download the Way of the Gopher VM

## Try it out; see if it all fits (install it, start it)

1. Install VirtualBox
2. Open VirtualBox and click 'Add', then navigate to the WayOfTheGopher.ovn file you downloaded

## Make a few practice swings (let's download a cat)

[VirtualBox]: https://www.virtualbox.org/wiki/Downloads
[windows]: https://download.virtualbox.org/virtualbox/6.1.16/VirtualBox-6.1.16-140961-Win.exe
[OS X]: https://download.virtualbox.org/virtualbox/6.1.16/VirtualBox-6.1.16-140961-OSX.dmg

---

# Download a cat!

    !go
    package main

    import (
        "io/ioutil"
        "net/http"
    )

    func main() {
        resp, _ := http.Get("https://i.imgur.com/7RRXZBR.jpg")
        body, _ := ioutil.ReadAll(resp.Body)
        ioutil.WriteFile("cat.jpg", body, 0644)
    }

## Place this code snippet in ~/go/src/github.com/over-codes/cat/main.go

## Run `go run github.com/over-codes/cat`

## Observe, cat!

---

# What does this mean?!?

## `package main`

- a package is a collection of code that gets run together, or shared.
- `main` has special meaning to programmers. It is the first place the computer looks when deciding to what to run; it is the main piece of code, and everything else just helps it

## `import ( ... )`

- `import` lets you pull in other packages and use them to help write your code
- Each package is either from the standard library, or is a URL that you can drop into a webbrowser

[standard library]: https://golang.org/pkg/

---

# What does this mean?!?

## `func main() {`

- programs are composed of functions (func, for short)
- Each function has 0 or more arguments (listed between the parenthesis) and 0 or more return values (which come after the close parenthesis)
- This function, the main function, has special meaning and always has 0 arguments and 0 return values

### Another example:

    !go
    func Get(url string) (resp *Response, err error) {

- This function is named Get, takes one argument, and returns two values
- We use this function, [http.Get] on the next line

[http.Get]: https://golang.org/src/net/http/client.go?s=15297:15345#L436

---

# What does this mean?!?

## `resp, _ := http.Get("https://i.imgur.com/7RRXZBR.jpg")`

- We declare and assign `resp` to the first return value of `http.Get`; we ignore the second value
- That second value has special meaning in Go convention; explained later. `_` means 'ignore'
- `http.Get` requires one argument; the URL to fetch
- `resp` is a variable; we can use it later

## body, _ := ioutil.ReadAll(resp.Body)

- We declare and assign `body` to the first return value of `ioutil.ReadAll`; we ignore the second value
- `ioutil.ReadAll` takes one argument; something to read from (the body of our HTTP response!)
- The returned value is a slice of bytes (`[]byte`), in this case representing a cat picture

---

# What does this mean?!?

## `ioutil.WriteFile("cat.jpg", body, 0644)`

- We ignore all return values of this function
- `ioutil.WriteFile` creates the file named in the first argument, and saves the contents of the second argument there. The third argument is called the mode
- `mode` is an old-school Unix term which is used to indicate the file is `readable`, `writable`, and/or `executable`
- `0644` means (read-write by the owner, and readable by everyone else)

## But what about `(`, `)`, `{`, and `}`?

- For every start, there is an end. In all cases, these parenthesis (`()`) must be matched, and the braces (`{}`) must be matched
- Things between parenthesis are arguments
- Things between braces are statements

---

# Let's make this better!

## Remember those `_` things meaning ignore this value? Never use them there.

- The second return value, or really any return value which is of type error, shouldn't be ignored
- It represents an error; so if your code didn't work, maybe you got an error
- For example, what if someone deleted that cat picture when we tried to get it? Sadness.

---

# Do this instead

    !go
    package main

    import (
        "io/ioutil"
        "log"
        "net/http"
    )

    func main() {
        resp, err := http.Get("https://i.imgur.com/7RRXZBR.jpg")
        if err != nil {
            log.Fatalf("Failed to fetch the image: %v", err)
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Failed to read all of the image: %v", err)
        }
        err = ioutil.WriteFile("cat.jpg", body, 0644)
        if err != nil {
            log.Fatalf("Failed to write the image to a file: %v", err)
        }
    }

---

# The new pattern

    !go
        resp, err := http.Get("https://i.imgur.com/7RRXZBR.jpg")
        if err != nil {
            log.Fatalf("Failed to fetch the image: %v", err)
        }

## If a function returns an error, check to see if the value is nil, if not, log it

- `if` is a Go keyword; `if something { doThis() }`. If the condition is true, run the code inside the following braces. `!=` stands for 'does not equal'
- `log.Fatalf` takes one or more arguments (a variable number of arguments, varargs)
- It doesn't return anything, but exits your program after printing something to the terminal
- The string as the first argument has special meaning for `log.Fatalf`; it can be formatted
- `%v` means 'use a nice representation of this value'

---

# Did you notice?

    !go
        err = ioutil.WriteFile("cat.jpg", body, 0644)
        if err != nil {
            log.Fatalf("Failed to write the image to a file: %v", err)
        }

## The `:=` changed to `=`

- This is because we are not declaring the `err` variable (it was already declared!)
- We are only assigning something to it

Consider this example:

    !go
        var path string // declare path to be a string
        path = "https://i.imgur.com/7RRXZBR.jpg" // assign a value
        resp, err := http.Get(path) // use it

- Recall that `http.Get` needed one string argument
- Programmers are lazy. `:=` lets you avoid writing the declaration, and infers the type of variable (i.e., string). It gets used a lot.

---

# Go forth and code!

## Before we go on, spend some time practicing with your new sword

## Make sure you feel comfortable with everything so far; it's OK if there are questions, as we did not cover everything

## Try to answer the questions yourself, with your Google-Fu, or ask on Discord

## Try this:

1. Download 4 cat pictures at once
2. Use `rand.Int` to select a random cat picture to download (you will need to use the `%` operator, mod; see [the list of operators], and [arithmetic operators])
3. Make the program fail; try to get each failure case (there were 3! the second is actually very hard to trigger; good luck!)
4. Join us on Discord!

[the list of operators]: https://golang.org/ref/spec#Operators_and_punctuation
[arithmetic operators]: https://golang.org/ref/spec#Arithmetic_operators