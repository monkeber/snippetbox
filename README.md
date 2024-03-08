# Snippetbox

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MariaDB](https://img.shields.io/badge/MariaDB-003545?style=for-the-badge&logo=mariadb&logoColor=white)

The project has been created while following the book "Let's Go" by Alex Edwards. Most of the exercises at the end of the book were implemented except the last one with functionality for password change.

The project does not use migration files, so all db setup was dumped in `schema.sql` file.

Tools used:

- MariaDB, v. 11.3.2
- Go, v. 1.22.1

## TLS Certificates

You also need to generate self-signed certificates, example:

```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

In my case this file was in:
```
/usr/lib/go/src/crypto/tls/generate_cert.go
```

## Default User

Credentials for the default user are:

```
Email: test@test.com
Password: testtest
```

## Build And Run

You can easily run the project with:

```
go run ./cmd/web
```

And access it on: `https://localhost:4000`
