Gotcha. Here’s the **step-by-step** to make Postman trust your self-signed cert and keep **SSL certificate verification ON**.

---

# A) Quick sanity check (what you have)

* Your Go server is running on **[https://localhost:3000](https://localhost:3000)** and uses **cert.pem** + **key.pem**.
* The cert’s **CN/SAN** should include `localhost`. If you used my earlier config, you’re fine.

---

# B) Add your certificate in Postman (Desktop app)

1. **Open Settings**

* Click the **gear icon** (top right) → **Settings**.

2. **Keep verification ON**

* In **General**, ensure **“SSL certificate verification”** is **ON** (blue).

3. **Add your CA/server certificate**

* Go to the **Certificates** tab.
* Scroll to **CA Certificates** (not “Client Certificates”).
* Click **“Add CA Certificate”**.
* **Choose file** → select your `cert.pem`.

  * Postman accepts `.pem/.crt/.cer`. Your `cert.pem` is OK.
* Toggle **Enabled** (if there’s a switch), then **Save**.
* **Restart Postman** (this helps it pick up new CA certs).

> Why CA Certificates? For a self-signed server cert, that file acts as its own issuer. Adding it here tells Postman to trust it.

4. **Send a request**

* Make a request to: `https://localhost:3000/teachers`
* You should now get a normal **200 OK** response (no SSL error).

---

# C) If you still see SSL errors

### 1) Check hostname/SAN

* The cert **must** have `localhost` in **CN** or **SAN**.
* If you’re unsure, re-generate with SAN:

  ```ini
  # openssl.cnf (minimal)
  [ req ]
  prompt = no
  distinguished_name = dn
  x509_extensions = v3_req

  [ dn ]
  CN = localhost

  [ v3_req ]
  subjectAltName = @alt
  [ alt ]
  DNS.1 = localhost
  IP.1 = 127.0.0.1
  ```

  ```bash
  openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout key.pem -out cert.pem -config openssl.cnf
  ```

  Then re-add **cert.pem** under **CA Certificates** in Postman, restart Postman, try again.

### 2) You accidentally added a **Client Certificate**

* That’s for **mutual TLS** (mTLS). You don’t need it unless your server requires client certs.
  If you added one under **Client Certificates**, remove it. Only use **CA Certificates** for trusting your server.

### 3) You’re using Postman **Web**

* TLS handshake happens locally; the web version can’t use your local certs. Use the **Postman Desktop app**.

### 4) Still stuck? Use a mini-CA (most reliable)

Create a root CA, then sign your server cert with it; import the **root CA** into Postman.

```bash
# 1) Root CA
openssl genrsa -out rootCA.key 4096
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1825 \
  -subj "/CN=My Local Dev CA" -out rootCA.crt

# 2) Server key + CSR
openssl genrsa -out key.pem 2048
openssl req -new -key key.pem -subj "/CN=localhost" -out server.csr

# 3) SAN extensions for server cert
cat > server.ext <<EOF
subjectAltName=DNS:localhost,IP:127.0.0.1
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth
EOF

# 4) Sign server cert with the root CA
openssl x509 -req -in server.csr -CA rootCA.crt -CAkey rootCA.key \
  -CAcreateserial -out cert.pem -days 825 -sha256 -extfile server.ext
```

* In **Postman → Settings → Certificates → CA Certificates → Add CA Certificate**, choose **`rootCA.crt`**.
* Restart Postman, hit `https://localhost:3000`.

---

# D) (Optional) Trust it system-wide on macOS

* Open **Keychain Access** → **login** → **Certificates**.
* Import `rootCA.crt` (or `cert.pem` if self-signed).
* Double-click → **Trust** → “**Always Trust**” for SSL.
* Some apps (Chrome) will now trust it; Postman may still use its own CA store—keep it added there too.

---

# E) When to use **Client Certificates** in Postman

Only if your server requires **mTLS** (client auth). Then in **Settings → Certificates → Client Certificates → Add Certificate**:

* **Host:** `localhost`
* **Port:** `3000`
* **CRT file:** your **client** certificate (not the server cert)
* **KEY file:** your client key

Your current Go server does **not** require client certs, so you don’t need this.

---

That’s it! Follow **B** (add CA certificate) and you can keep **SSL verification ON** without errors. If you want, I can also give you a tiny script to **verify your cert fields** (CN/SAN) so you know the Postman trust issue isn’t from the certificate content.
