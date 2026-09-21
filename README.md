-----------

<div align="center">

![](https://img.shields.io/badge/Go-1.25.3-blue?style=plastic&logo=go)&nbsp; 
![](https://img.shields.io/badge/Git-2.32.1-red?style=plastic&logo=git)&nbsp;

</div>

----------

# platform-dev-tool-kit
a cli tool for developers prepared by platform engineer.
> Jump into the world where you can help the dev team without being a gate keeper in the process.

# Supported Architecture

- MacOs (Apple Silicon)
- Linux 
- Windows (not tested yet)

# Get started

### Linux/Mac users -
```bash
chmod +x install.sh;
./install.sh
```

# for windows uses
- generate binary
```bash
GOOS=windows GOARCH=amd64 go build -o platform.exe
```

- move to this directory `C:\Users\<username>\bin\`

# Features
0. Configure - `platform configure` 
    - asks for API token and ticket platform url and configurs the binary to use that late. like aws configure works.
1. Create branch as per ticket convention - `platform start ticket EP-135`
    1.1. fetch youtrack ticket id, description 
        1.2. create state file in the repo 
            1.3. auto ignore the state file in the repo


> ashraf-minhaj was here