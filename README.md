## Simple crypto library 

### Usage

```go
func main() {
    urPasswordForEncryption := "password"
    c := crypto.New(urPasswordForEncryption)

    // encrypt file nad ouptut encrypted file with .enc postfix
    if err := c.EncryptFile("hello.txt", urPasswordForEncryption); err != nil {
    	panic(err)
    }

    // decrypt file nad ouptut decrypted file with .dec postfix
    if err := c.DecryptFile("hello.txt.enc", urPasswordForEncryption); err != nil {
    	panic(err)
    }

    // encrypt text and return base64 encoded string
    encodeText, err := c.EncryptText("hello", urPasswordForEncryption)
    if err != nil {
    	panic(err)
    }
    log.Print(encodeText) // base64 encoded string

	// decrypt text from base64 encoded string
    decodeText, dErr := c.DecryptText(encodeText, urPasswordForEncryption)
    if err != nil {
    	panic(dErr)
    }

    log.Print(decodeText) // hello
}
```

![mem](./public/mem.png)
