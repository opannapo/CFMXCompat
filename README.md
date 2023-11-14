# gocfmx
**CFMX_COMPACT** algorithm with an encrypt function to convert to **hexadecimal** (_Hex_)


**Example :**
```
cfMx := gocfmx.NewCFmxCompat("Key123456789")

encResult := cfMx.EncryptEncodingHex("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
fmt.Printf("Encrypt : %s \n",enc)

decryptResult, _ := cfMx.DecryptEncodingHex("77789CBA0A079E5554084A1BB8800091BED999A5C0F35AD258AF53DFBA71A1EA885A439A93A26ED69C7A968E8E6A75BA5D51DF9988A28B48")
fmt.Printf("Decrypt : %s \n",decryptResult)
```

**Result :**
```
Encrypt : 77789CBA0A079E5554084A1BB8800091BED999A5C0F35AD258AF53DFBA71A1EA885A439A93A26ED69C7A968E8E6A75BA5D51DF9988A28B48
Decrypt : Lorem ipsum dolor sit amet, consectetur adipiscing elit.
```
