# WebDAV

Setup your WebDAV server quickly.

## Environments as cli options

```
dir := getEnv("ROOT", ".") // root of webdav
httpPort := getEnv("HTTP_ADDR", ":80")
httpsPort := getEnv("HTTPS_ADDR", ":443")
serveSecure := getEnv("HTTPS_ENABLE", "") != "" // enable https
certKey := getEnv("CERT_KEY", "cert.pem")
pubKey := getEnv("PUB_KEY", "key.pem")
authUsers := getEnv("AUTH_PATH", "auth.json") // HTTP basic auth table in name-password format
```
