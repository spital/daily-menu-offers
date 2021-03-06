# daily-menu-offers

Write an app to find out what's cooking today.

Every pub presents information about its daily or weekly menu somewhere on the web. Your job is to find out and present information about what each of the restaurants below is cooking today. I like toppings, so I'm interested in those too. Just the whole daily menu... It's good to know what's cooking when you don't know what you're craving.

## CO SE DNES VARI?

Napis aplikaci, ktera zjisti co se prave dnes vari.

Kazda hospoda nekde na webu prezentuje informace o svem dennim, ci tydenim
mennu. Tvym ukolem je zjistit a prezentovat informaci o tom, co prave dnes
kazda z nize uvedenych restauraci vari. Mam rad polevky, takze me zajimaji i ty.
Proste cela denni nabidka...
Dobre vedet co kde prave vari, kdyz nevis na co mas chut.

PIVNICE U CAPA
https://www.pivnice-ucapa.cz/denni-menu.php

SUZIES STEAK PUB
http://www.suzies.cz/poledni-menu

VERONI CAFE
https://www.menicka.cz/4921-veroni-coffee--chocolate.html


## Install Golang

For fedora 33
```bash
sudo dnf -y install golang
```
Or check https://golang.org/doc/install

## Test

*Note* test only mutex/waitgroup/channel of string; web parsing test is not included as I did not want to upload pieces of websites

```bash
go test
```

## Build and run app

```bash
go run ./main.go

# or

go build && ./daily-menu-offers
```
