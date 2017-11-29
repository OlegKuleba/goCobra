# goCobra

It's a GO-application, which uses Cobra technology

### USAGE
All of commands are used with go run main.go {command} [arg1, arg2, ..., argN]

### COMMANDS' FORMAT
#### addContact
addContact phoneNumber name city street building [apartment]<br>
ex: go run main.go addContact +380994445666 Taras Dnipro Titova 26a 45
#### findAll
findAll<br>
ex: go run main.go findAll
#### findByNumber
findByNumber phoneNumber<br>
ex: go run main.go findByNumber +380994445666
#### editContact
editContact phoneNumber name city street building [apartment]<br>
ex: go run main.go editContact +380994445666 Taras Kyiv Shevchenka 15
#### deleteContact
deleteContact phoneNumber<br>
ex: go run main.go deleteContact +380994445666