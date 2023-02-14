# golang-learning-cloud-crew

A repo to play around with golang

## Install GoLang

### Linux
1. Copy the download link from the official golang website https://go.dev/doc/install

2. Download the tarball
   e.g. 
   ```
   https://go.dev/dl/go1.20.linux-amd64.tar.gz
   ```

3. Delete prvious versions of the tarball and extract this new one
   ```
   rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
   ```

4. Add Go binary to path
   ```
   export PATH=$PATH:/usr/local/go/bin
   ```

5. Check that installation was successful
   ```
   go version
   ```
