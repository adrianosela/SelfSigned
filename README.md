# SelfSigned
Self-Signed CA SSL certificate (and private key) generator. 

## Usage:

Build The project binary:

```
$ make
```

Run the binary with the desired CN on the CA Certificate as the first argument:

```
$ ./self-signed adrianosela.com
```

The certificate will be saved as cert.pem and the private key (which you will use to sign certificates on behalf of the new CA) as key.pem:

```
$ cat key.pem
-----BEGIN PRIVATE KEY-----
<<<<<<< raw bytes >>>>>>>>
-----END PRIVATE KEY-----

$ cat cert.pem
-----BEGIN CERTIFICATE-----
<<<<<<< raw bytes >>>>>>>>
-----END CERTIFICATE-----
```

You may use the following OpenSSL command to inspect your certificate:

```
$ openssl x509 -in cert.pem -text -noout
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            7d:59:ff:af:e2:68:db:c1
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=CA, O=AdrianoSela Inc., CN=adrianosela.com
        Validity
            Not Before: Dec 12 01:38:17 2018 GMT
            Not After : Dec  9 01:38:17 2028 GMT
        Subject: C=CA, O=AdrianoSela Inc., CN=adrianosela.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
            RSA Public Key: (2048 bit)
                Modulus (2048 bit):
                    00:d6:7f:03:fa:fe:84:6c:3d:59:bd:2e:5b:93:3f:
                    bf:89:d0:f7:e8:ba:6c:ca:60:ed:b8:75:97:8b:e8:
                    45:88:b8:81:b9:d6:cf:53:90:31:95:bc:15:76:8e:
                    1f:ab:cf:3e:e4:a6:6f:a9:ab:17:1d:02:98:ee:b4:
                    a2:f3:9c:76:a9:bd:a7:2c:25:a0:da:ec:a7:34:29:
                    a7:ee:da:ca:28:c0:af:47:5d:a2:dc:5c:f3:55:4d:
                    10:40:6f:dc:98:0d:46:88:dd:af:39:b5:ab:e7:e6:
                    87:8f:13:92:c8:04:fd:35:58:00:6b:e0:14:01:bc:
                    05:7a:14:23:e5:e5:6c:ee:a6:81:4d:0c:7f:6f:bb:
                    dc:e5:c3:45:d9:1c:43:dd:dc:4d:d8:57:fe:57:a0:
                    2b:90:b6:ed:a0:83:4b:46:18:12:61:d5:65:83:71:
                    46:07:9c:3f:02:48:49:7f:8b:40:76:34:b4:92:2f:
                    49:5b:f1:b4:40:a6:8d:75:71:36:a0:c2:0b:3a:54:
                    4d:7f:ce:db:6c:b9:5b:c8:eb:ce:06:10:46:3e:22:
                    c8:a4:c0:e7:ad:45:d2:9d:25:89:75:32:72:e2:7f:
                    a5:01:80:d6:f0:f1:2b:98:94:ce:56:57:6d:5d:d1:
                    fa:7f:b7:50:f6:b1:9b:6b:ed:e4:dd:85:0c:97:c7:
                    b5:49
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment, Certificate Sign
            X509v3 Extended Key Usage:
                TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:TRUE
    Signature Algorithm: sha256WithRSAEncryption
        3d:dc:ca:22:ee:f5:1b:d8:c3:4b:b8:89:2b:a9:c2:ac:db:9b:
        1f:7c:b6:47:4d:84:43:d7:c6:b4:51:23:6a:ae:1b:bc:47:b9:
        4a:13:c6:90:00:26:d2:c1:c0:aa:80:41:4a:d3:33:f3:fd:6e:
        15:22:c2:df:3b:e9:76:fe:e0:ce:75:e7:80:4b:00:a7:7a:50:
        9d:2d:12:82:e8:51:9e:27:8a:dc:77:98:96:5e:88:ed:c4:90:
        ce:f1:55:1c:a4:bc:9a:db:da:54:07:c8:50:36:d8:39:64:d2:
        70:ec:44:f5:e1:b8:b4:13:10:0d:1c:a6:bf:b9:3e:b6:9b:5e:
        cd:a3:f3:dd:39:6c:f1:e0:1f:79:30:38:73:e5:98:13:24:01:
        b9:2c:f4:1e:18:97:61:f5:d7:b7:ae:be:7b:bf:fc:35:bb:2b:
        f8:10:fe:ff:29:d7:b2:c9:0b:b5:ad:d8:f7:0f:ba:2a:47:28:
        dd:c7:8b:f8:a7:f7:31:43:55:be:10:48:a4:bd:84:16:c9:e5:
        1d:8c:08:0f:01:7c:80:35:31:f9:a1:c9:02:e5:9c:8b:34:5e:
        5a:c3:c5:f7:a3:e7:12:2f:b8:7c:89:b0:38:7d:67:43:0e:b2:
        93:aa:f7:e6:c8:96:64:f9:93:a7:e1:f7:aa:40:b0:b7:5b:43:
        ed:3b:cd:78
```