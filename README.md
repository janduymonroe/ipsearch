# Project Name

Welcome to Project Name! This project is a tool built in Go that allows you to retrieve the IP address or server name of a website. It provides two commands, `ip` and `servers`, along with a `--host` parameter for specifying the website. This document will guide you on how to build, run, and utilize the commands.

## Prerequisites

Before running this project, ensure that you have the following prerequisites installed:

- [Go](https://golang.org/dl/) (version 1.19 or later)

## Installation

To install and set up the project, follow these steps:

1. Clone the repository:

```bash
    git clone https://github.com/janduymonroe/ipsearch.git
```

2. Navigate to the project directory:

```bash
    cd project-name
```

## Install the dependencies:

```bash
go get github.com/urfave/cli v1.22.14
```

## Usage
This project provides two main commands: ip and servers. These commands allow you to retrieve the IP address or server name of a given website.

## Command: ip
The ip command returns the IP address of the specified website. To use this command, follow the syntax below:

```bash
go run main.go ip --host <website>
```

Replace <website> with the URL of the website for which you want to retrieve the IP address. For example:

```bash
go run main.go ip --host www.example.com
```

The command will output the IP address of the specified website.

## Command: servers
The servers command returns the server name of the specified website. To use this command, follow the syntax below:

```bash
go run main.go servers --host <website>
```

Replace <website> with the URL of the website for which you want to retrieve the server name. For example:

```bash
go run main.go servers --host www.example.com
```

The command will output the server name of the specified website.

## Customization
You can modify the default commands or add new functionalities to suit your needs. To do so, open the main.go file and make the necessary changes.

## License
This project is licensed under the MIT License. Feel free to use, modify, and distribute it as per the terms of the license.

## Contribution
Contributions to this project are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## Contact
If you have any questions or need assistance, feel free to contact the project maintainer at your-email@example.com.

Thank you for using Project Name! We hope you find it helpful in retrieving IP addresses and server names for websites.

