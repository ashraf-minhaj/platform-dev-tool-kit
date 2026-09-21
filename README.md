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

### for windows users -
- generate binary
```bash
GOOS=windows GOARCH=amd64 go build -o platform.exe
```

- move to this directory `C:\Users\<username>\bin\`

# Features

### 0. Configuration

**`platform configure`**

* Configure YouTrack URL
* Configure API token
* Store configuration globally for future commands
* Works similar to `aws configure`

### 1. Ticket-Based Development

**`platform start ticket EP-135`**

#### 1.1 Fetch Ticket

* Fetch ticket details from YouTrack
* Retrieve ticket ID, summary, and description

#### 1.2 Project State

* Create `.platformState.json` in the repository
* Store the current ticket information
* Automatically add `.platformState.json` to `.gitignore`
* Create `.gitignore` if it does not exist

#### 1.3 Create Feature Branch

* Generate a branch name from the ticket ID and title
* Create and checkout the feature branch

Example for ticket: EPD-300: Implement New Authentication System

Command - `platform start ticket EPD-300`, and a new branch named 
```text
→ feature/EPD-300-implement-new-authentication-system
```

is created.

> More to come. This is under development. Feel free to contribute. How do you do it? If you are a DevOps you should already know it. 

---

<div align="center">

[ashraf-minhaj](https://www.linkedin.com/in/ashraf-minhaj) was here

</div>