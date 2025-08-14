 Master The Go Programming Language: Elevate Your Skills!

 Go Bootcamp with gRPC and Protocol Buffers

Welcome to the ultimate journey into Go programming! Whether you're a beginner aiming to dive headfirst into software development or an experienced coder looking to sharpen your skills, this comprehensive course is tailored just for you.

Important Note
🧠 Why Do I Get "main redeclared" Errors in Go? Understanding package main, main() Functions, and Folder Structure in Your Go Course
If you're seeing an error like:

main redeclared in this block
Or you're wondering why you can’t go run a file after changing the package name, you’re not alone. Many students encounter this while following along with the course, so let’s break it down step-by-step and clarify how Go handles files, packages, and the main function.

✅ The Setup in This Course
In this course, we build understanding one file at a time — first hello_world.go, then data_types.go, variables.go, and so on.

But only one file exists in the main folder at a time.

Once a concept is covered in a lecture, that file is either:

🔄 Renamed (e.g., to .txt), so Go won’t try to compile it.

📁 Moved into a subfolder like main/, and its package line is updated to match the folder (package main).

This is done intentionally to keep the learning experience clean and prevent errors like multiple main() declarations.

🧨 Why Does This Error Happen?
Let’s say you have two files in the same folder:

random.go

package main
 
func main() {
    fmt.Println("Random code")
}
variables.go

package main
 
func main() {
    fmt.Println("Variables code")
}
You will get an error:

main redeclared in this block
This is because Go allows only ONE main() function per package, and by default, all .go files in a folder belong to the same package (unless explicitly specified otherwise).

💡 What is package main?
In Go, if you're writing an executable program, it must:

Use package main

Have a main() function as the entry point

If you try to run a file without package main or without a main() function, you’ll get an error like:

go run: cannot run non-main package
💡 Then What is package main, package utils, etc.?
These are library-style packages — used for reusable code. They don’t need a main() function, and they’re not meant to be executed directly with go run.

That’s why, in some lectures, I change the file from:

package main
to:

package main
Before archiving or moving it — this tells Go, "Hey, this is not a program to run anymore."

🛠️ Your Options
Here’s how you can avoid or fix these errors:

✅ Option 1: Follow the folder and naming structure exactly as I am doing in the lectures
✅ Option 2: Keep Only One main() at a Time
If you want to keep everything in one folder (e.g., during early learning), just make sure only one file has a main() function.

You can temporarily rename other .go files to .txt or comment out extra main() functions.

✅ Option 3: Use Subfolders for Different Programs
Structure your code like this:

/Go_course/
  /random/
    random.go   --> package main (has its own main())
  /variables/
    variables.go --> package main (has its own main())
Now you can cd into each folder and run:

go run random.go
Each file is part of a different Go program now. No conflicts!

✅ Option 4: Combine Logic Into One File
Instead of having two main() functions, split functionality into named functions:

package main
 
func main() {
    sayHello()
    showVariables()
}
 
func sayHello() {
    fmt.Println("Hello from Random")
}
 
func showVariables() {
    fmt.Println("Hello from Variables")
}
This is good for practice and helps organize your code better.

🤔 Why Do I Rename Files or Change the Package?
When you see me renaming files or changing package main to package main, it’s to avoid compilation errors. This way, I can keep the previous lecture’s code as a reference without Go thinking I’m trying to run multiple programs at once.

⚠️ Important Rules to Remember
✅ Allowed                                                                             |                 ❌ Not Allowed

----------------------------------------------------------------------------------------------------------------------

One main() in a package Multiple                                   |                 main() functions in the same package

Files in different folders, each with package main       |                  Two main() files in the same folder

go run file.go with package main                              | go run on a file with a different package name

Archiving old .go files by changing package or extension |  Leaving many .go files with main() in one folder

🔁 FAQ Summary
Q: Can I have multiple .go files with main()?
A: Yes, but only if they’re in separate folders or packages.

Q: Why change the package name to main?
A: To indicate that the file is now part of a library, not meant to be run.

Q: Why does go run give an error after changing the package name?
A: Because only package main can be run as an executable.

Q: Why does VS Code show squiggly lines when I have multiple files?
A: Because it detects multiple main() functions or duplicate declarations.



💡 Why I’m Only Changing the Package Name (Not the main() Function)

You might have noticed that inside the main, intermediate, and advanced folders, many files still have a main() function — and that’s completely okay. I don’t rename the main() function in those files because we’re not trying to run them directly while they’re sitting in those folders. Each of those folders is like a “storage area” for code examples we’ve already covered. If you tried to run any of those folders directly, you’d get an error — because Go doesn’t allow multiple files with main() functions in the same package (i.e., folder) to be compiled together.

Instead of spending time removing or renaming every main() function, I simply change the package name to match the folder (package main, package intermediate, etc.). That way, the files are still saved for reference, but Go won’t treat them as executable programs anymore — it will treat them like regular library code that’s not meant to be run directly.

Later, if you want to run any of those example files again, just move the file out into the main folder, rename its package to main, and make sure there’s no other file with a main() function in that folder. Then you can run it normally using:

go run your_file.go
This setup helps us keep things clean, focused, and organized by topic, without getting stuck on Go’s restrictions around main() functions.



As we progress into the course and as you grow more comfortable, you’ll learn how to structure projects more professionally (and there is an exclusive coverage on this as well, later on), but right now, the goal is understanding Go fundamentals — not full-scale architecture.


# Simple API Server with HTTP/2 and Mutual TLS (mTLS) in Go

This project is a **Go HTTP/2 API server** that uses **TLS** and **mutual TLS authentication (mTLS)** for secure communication between clients and the server.

It includes:

* Two API endpoints: `/orders` and `/users`
* TLS 1.2+ encryption
* Client certificate validation (mTLS)
* HTTP/2 support

---

## 📂 Project Structure

```
.
├── server.go          # Main server code
├── cert.pem           # Server certificate
├── key.pem            # Server private key
├── openssl.cnf        # OpenSSL configuration for certificate generation
├── go.mod             # Go module definition
├── go.sum             # Go dependencies
```

---

## ⚙️ Prerequisites

Make sure you have:

* **Go** (>= 1.23.x) installed
  [Download Go](https://go.dev/dl/)
* **OpenSSL** installed
  On macOS:

  ```sh
  brew install openssl
  ```

  On Ubuntu/Debian:

  ```sh
  sudo apt update && sudo apt install openssl
  ```

  On Windows:
  [Download OpenSSL for Windows](https://slproweb.com/products/Win32OpenSSL.html)

---

## 🛠 Step 1: Clone the Project

```sh
git clone https://github.com/your-username/simple_api.git
cd simple_api
```

---

## 🔑 Step 2: Create TLS Certificates

This server requires **server certificates** and a **client CA** for mTLS.
You can either:

* Use the provided `openssl.cnf` to create certificates, OR
* Use a quick one-line OpenSSL command.

### **Option 1: Using Configuration File**

1. Create an OpenSSL config file (`openssl.cnf`) with this content:

   ```ini
   [req]
   default_bits       = 2048
   distinguished_name = req_distinguished_name
   req_extensions     = req_ext
   prompt             = no

   [req_distinguished_name]
   C  = US
   ST = State
   L  = City
   O  = Organization
   OU = Organizational Unit
   CN = localhost

   [req_ext]
   subjectAltName = @alt_names

   [alt_names]
   DNS.1 = localhost
   DNS.2 = 127.0.0.1
   ```
2. Generate server certificate and key:

   ```sh
   openssl req -x509 -nodes -days 365 \
     -newkey rsa:2048 \
     -keyout key.pem \
     -out cert.pem \
     -config openssl.cnf
   ```

### **Option 2: Quick Command**

```sh
openssl req -x509 -newkey rsa:2048 -nodes \
  -keyout key.pem -out cert.pem -days 365
```

When prompted, enter details:

```
Country Name (2 letter code) [AU]: AU
State or Province Name (full name) [Some-State]: Non Existent
Locality Name (eg, city) []: Random
Organization Name (eg, company) [Internet Widgits Pty Ltd]: API Inc
Organizational Unit Name (eg, section) []: API Inc
Common Name (e.g. server FQDN or YOUR name) []: API Inc
Email Address []: test@test.com
```

---

## 📜 Step 3: Install Dependencies

```sh
go mod tidy
```

---

## 🚀 Step 4: Run the Server

```sh
go run server.go
```

You should see:

```
Server is running on port: 3000
```

---

## 🌐 Step 5: Testing the Server

Since the server enforces **mTLS**, clients must present a valid certificate signed by the CA.
Example `curl` command:

```sh
curl --http2 -v https://localhost:3000/orders \
  --cert client-cert.pem \
  --key client-key.pem \
  --cacert cert.pem
```

---

## 📌 Important Notes about Go `main()` Functions

This project has **only one** `main()` function in `server.go`.
If you add more files, remember:

* Only one `main()` per package
* Keep different Go programs in separate folders to avoid `main redeclared` errors

For details, see the [Go Bootcamp Notes](#) in this README for an explanation.

Simple API with mTLS in Go
This project demonstrates a simple Go HTTP server with mutual TLS (mTLS) authentication. The server exposes two endpoints, /orders and /users, and enforces mTLS to ensure secure communication. The server uses HTTP/2 and logs request details, including the HTTP version and TLS version.
Prerequisites
To run this project, ensure you have the following installed:

Go: Version 1.23.3 or later
OpenSSL: For generating TLS certificates
Postman: For testing the API with mTLS
curl: For testing via the terminal
A modern web browser (e.g., Chrome) for testing

Project Structure
simple_api/
├── cert.pem         # Server certificate
├── key.pem         # Server private key
├── server.go       # Main Go server code
├── go.mod          # Go module file
├── go.sum          # Go module dependencies checksum
└── openssl.cnf     # OpenSSL configuration file

Setup Instructions
Step 1: Clone or Create the Project

Create a directory for the project:mkdir simple_api
cd simple_api


Initialize a Go module:go mod init github.com/ankitakapadiya/simple_api


Create or copy the server.go, go.mod, go.sum, and openssl.cnf files into the simple_api directory.

Step 2: Install Dependencies
Run the following command to download the required Go dependencies:
go mod tidy

This will install the golang.org/x/net package specified in go.mod.
Step 3: Generate TLS Certificates
The server requires a certificate (cert.pem) and private key (key.pem) for TLS. Follow these steps to generate them using OpenSSL:

Ensure you have the openssl.cnf file in the project directory with the following content:
[req]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext
prompt             = no

[req_distinguished_name]
C  = US
ST = State
L  = City
O  = Organization
OU = Organizational Unit
CN = localhost

[req_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
DNS.2 = 127.0.0.1


Generate the server certificate and key:
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365 -config openssl.cnf


This command creates a self-signed certificate valid for 365 days.
The -nodes flag ensures the private key is not encrypted.
The -config openssl.cnf flag uses the provided configuration to set localhost and 127.0.0.1 as valid domains.


(Optional) Generate a client certificate for mTLS testing:
openssl req -x509 -newkey rsa:2048 -nodes -keyout client-key.pem -out client-cert.pem -days 365 -config openssl.cnf


This creates client-cert.pem and client-key.pem for testing mTLS.
Use the same openssl.cnf to ensure compatibility with the server certificate.



Step 4: Run the Server

Ensure cert.pem and key.pem are in the simple_api directory.
Start the server:go run server.go


The server will run on https://localhost:3000.
It listens for requests on /orders and /users endpoints.
The server enforces mTLS, so clients must provide a valid certificate.



Testing the API
Testing with Postman
Postman supports mTLS configuration for testing secure APIs. Follow these steps:

Open Postman and create a new request.
Set the request URL to:
https://localhost:3000/orders or https://localhost:3000/users


Configure mTLS in Postman:
Go to Settings > Certificates tab.
Click Add Certificate.
Enter the following details:
Host: localhost:3000
CRT file: Select client-cert.pem
KEY file: Select client-key.pem
Passphrase: Leave blank (since -nodes was used).


Click Add.


Send the request:
Method: GET
URL: https://localhost:3000/orders
You should receive a response: Handling incoming orders.
Similarly, test https://localhost:3000/users to get Handling users.


Troubleshooting:
If you see a certificate error, ensure the client certificate is valid and matches the server's expected CA (cert.pem).
Disable SSL certificate verification in Postman settings if you encounter issues with self-signed certificates:
Go to Settings > General > Turn off SSL certificate verification.





Testing in Chrome
Browsers do not support mTLS natively, so you may encounter issues due to the self-signed certificate or mTLS requirements. To test in Chrome:

Open Chrome and navigate to https://localhost:3000/orders.
You may see a "Your connection is not private" warning due to the self-signed certificate.
Click Advanced > Proceed to localhost (unsafe) to bypass the warning.
Note: Chrome cannot provide a client certificate for mTLS, so the server will reject the request with a 400 Bad Request or similar error unless mTLS is disabled on the server (not recommended).
For proper mTLS testing, use Postman or curl instead.

Testing with curl in Terminal
Use curl to test the API with mTLS:

Run the following command to test the /orders endpoint:curl --cert client-cert.pem --key client-key.pem --cacert cert.pem https://localhost:3000/orders


Expected output: Handling incoming orders


Test the /users endpoint:curl --cert client-cert.pem --key client-key.pem --cacert cert.pem https://localhost:3000/users


Expected output: Handling users


Troubleshooting:
If you get a curl: (60) SSL certificate problem, ensure the --cacert cert.pem flag is used to trust the server's self-signed certificate.
If mTLS fails, verify that client-cert.pem and client-key.pem are valid and match the server's CA (cert.pem).



Endpoints

GET /orders: Returns Handling incoming orders.
GET /users: Returns Handling users.

Notes

The server uses HTTP/2 and requires TLS 1.2 or higher.
mTLS is enforced, so all clients must provide a valid client certificate signed by the CA in cert.pem.
Logs in the server console display the HTTP version and TLS version of incoming requests.
For production, replace self-signed certificates with ones issued by a trusted CA.

Troubleshooting

"main redeclared" error: Ensure only one main() function exists in the package main. If you have multiple .go files with main() functions in the same directory, move them to separate directories or rename them (e.g., to .txt).
Certificate errors: Verify that cert.pem, key.pem, client-cert.pem, and client-key.pem are in the project directory and correctly referenced.
Port conflict: If port 3000 is in use, change the port variable in server.go to another value (e.g., 8080).

Learning Resources

Go Documentation
OpenSSL Documentation
Postman mTLS Guide
HTTP/2 in Go

Happy coding! 🚀