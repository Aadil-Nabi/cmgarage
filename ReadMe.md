# cmgarage
# Welcome to cmgarage app 👋
This is a small app to check the overall status of the below services on Thales CipherTrust Manager
- Password Validation
- Disk Encryption
- Backup Status
- Cluster Status

# Installation ✔
Make sure you have go installed on your workstation and you can download the code and use it directly on any IDE

# Usage ▶
* Create a config.yaml file, example as below
```bash
env: "dev"
cm_secret:
  base_url: "https://192.168.238.129/api/"
  version: "v1"
  cm_user: "admin"
  cm_password: "ADggebwl@1234#^&"
  encryption_key: "cm17"
```
To run the program you need to execute the run command in the below format
```bash
go run cmd/cmgarage/main.go -configfile config.yaml
```
Build the program into an executable file for your specific platform, Example below for Windows
```bash
go build -o cmgarage.exe cmd/cmgarage/main.go -configfile config.yaml
```

# Contributing 🤝
Contributions, issues and feature requests are welcome, but it is paused for now.
Feel free to check issues page if you want to contribute in future
