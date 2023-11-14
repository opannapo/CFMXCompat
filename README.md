# gocfmx
**CFMX_COMPACT** algorithm with an encrypt function to convert to **hexadecimal** (_Hex_)


**Example :**
```
cfMx := gocfmx.NewCFmxCompat("Key123456789")

encResult := cfMx.EncryptEncodingHex("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
fmt.Printf("Encrypt : %s \n",enc)

decryptResult, _ := cfMx.DecryptEncodingHex("77789cba0a079e5554084a1bb8800091bed999a5c0f35ad258af53dfba71a1ea885a439a93a26ed69c7a968e8e6a75ba5d51df9988a28b48")
fmt.Printf("Decrypt : %s \n",decryptResult)
```

**Result :**
```
Encrypt : 77789cba0a079e5554084a1bb8800091bed999a5c0f35ad258af53dfba71a1ea885a439a93a26ed69c7a968e8e6a75ba5d51df9988a28b48
Decrypt : Lorem ipsum dolor sit amet, consectetur adipiscing elit.
```
