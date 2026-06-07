![Development Status](https://img.shields.io/badge/status-active_development-green)

## Status: Development ##
* I have improved the code beyond the architecture and initial structure. I have added recursion to folder search and path information (following the clean logic of Unix-like systems).
* I introduced a new item in the app's logic: auth.go. This actually can hold multiple authentication and/or permission definion functions for users + some other access rules for different data points. Currently I defined here the `IsAdmin` function which should check if the user is an admin when they try to execute write actions and/or go do in-depth searches in the filesystem.
* Fixed some bugs, did some tweaks, and implemeneted additional helpers.
* Overall, it is a functional build / application.
* Gave the GitHub Actions a go (pun intended), and it seem to perform nicely as well!


## Motivation ##
The project is intended to solve real-world operational challenges, specifically focused on efficient system log management and task fluidization.

This content and application have the sole purpose of learning and practing idiomatic Golang development + system optimization, and task fluidization. As such, I used a free-open-source license for maxium flexibility and transparency.

## About ##
`syslog-utils-go` is a go system log management utility; it is desinged to address the need of managing system logs from a given system. The main logic includes the following operations: *directory management*, *log operations*, and *operation reliability*.

Due to simple yet powerful file-handling Go methods, the expectation for the following utility is to be able to perform the following actions:
1. Create log files
2. Read log files
3. Write to log files
4. Create directories
5. Change the working directory
6. Handle data visualization within the filesystem tree.
7. Proper error handling (at every step)
