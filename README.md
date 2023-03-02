# go-swifdog

This library provides a Golang API client for the Swifdog API. It is mainly maintained by the core team.

## Examples

## Clients

### Basic Client

```go
client, _ := NewBasicClient("email", "password")
```

### Bearer Token Client

```go
client, _ := NewBearerTokenClient("oylI1I4N7QDtf8WuI9UItKACpnwaI69MQZhKELeRzJfzovMmZVdvFKzrOCCr9S7j")
```

### Print the current account details

```go
account, err := client.GetAccount()
if err != nil {
	log.Fatal(err)
}

log.Println(account)
```

### List account tokens

```go
tokens, err := client.ListAccountTokens()
if err != nil {
    log.Fatal(err)
}

log.Println(tokens)
```

### Create a new account token

```go
str := "test by api client"
newToken, err := client.CreateAccountToken(&str) // <-- description can be null!
if err != nil {
    log.Fatal(err)
}

log.Println(newToken)
```

### Delete an existing account token by its identifier

```go
err = client.DeleteAccountTokenById("5279f81f-b50b-46e7-a4f5-7a074eb2f1e1")
if err != nil {
    log.Fatal(err)
}
```
## To-Do's

-[ ] Account (Get)
-[ ] Account Tokens (List, Create, Read, Write, Delete)
-[ ] Projects (List, Create, Patch, Read, Write, Delete)
-[ ] Volumes (List, Create, Patch, Read, Write, Delete)
-[ ] Packets (List, Create, Patch, Read, Write, Delete)
-[ ] IngressRules (List, Create, Patch, Read, Write, Delete)