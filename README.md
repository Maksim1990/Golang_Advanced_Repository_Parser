# Golang_Advanced_Repository_Parser
Golang advanced parser for directories and/or files

### How To Insatall

### 1) Clone current repository:
```
git clone https://github.com/Maksim1990/Golang_Advanced_Repository_Parser.git [APP_NAME]
```
### 2) Navigate to the clonned derictory
```
cd [APP_NAME]
```
## THIS IS ALL!!!

### How to use with different options

### 1) For simple displaying of all `directories and subdirectories` run following command
```
go run main.go .
```
![Mockup for feature A](https://github.com/Maksim1990/Golang_Advanced_Repository_Parser/blob/master/img/go1.PNG?raw=true)

### 2) For simple displaying of all `directories and subdirectories AND INCLUDED FILES` run following command
```
go run main.go . -f
```
![Mockup for feature A](https://github.com/Maksim1990/Golang_Advanced_Repository_Parser/blob/master/img/go1_2.PNG?raw=true)

### 3) For simple displaying of all `HIDDEN directories and subdirectories` run following command
```
go run main.go . -h
```
![Mockup for feature A](https://github.com/Maksim1990/Golang_Advanced_Repository_Parser/blob/master/img/go1_3.PNG?raw=true)

### 4) For simple displaying of all `HIDDEN directories and subdirectories AND INCLUDED FILES` run following command
```
go run main.go . -h
```
![Mockup for feature A](https://github.com/Maksim1990/Golang_Advanced_Repository_Parser/blob/master/img/go1_4.PNG?raw=true)

## Unit Tests

### FOR START UNIT TESTS RUN FOLLOWING COMMAND
```
go test -v
```
