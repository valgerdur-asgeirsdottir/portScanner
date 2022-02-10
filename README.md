Concurrent Port Scanner - PROGRAMMING ASSIGNMENT 2
Group A12 - Katrín Viktoría Hjartardóttir and Valgerður Ásgeirsdóttir

- Files

* scanner.go
This file contains the program using three functions: incrAddress, checkConnection and main.
Explainations of the functions and data flow can be found in the code comments of this file.
* go.mod
This file creates a new go module, initializing go for this project.
* .gitignore
Since git was used for collaboration in this project, this file contains the names all files and directories to be ignored when commiting that are relevant for a go project.

- Explanations and reasoning

* How do you determine the number of goroutines scanning hosts concurrently?

We decided to have a maximum of 1024 goroutines running at a time since there can only be a certain number of file descriptors that are free to run on the computer at each time. When a port is scanned, a new file descriptor is opened and therefore only this limited number of ports can be scanned at a time.

* How and under what conditions do you start the goroutines?

The goroutines are started after an attemt to write to a channel (file scanner.go, function main, line 110) with a maximum of 1024 inputs at a time (because of the channel size set in file scanner.go, function main, line 105). If the channel is full at the time the program waits until another goroutine has finished (reading from the channel in file scanner.go, function checkConnection, line 38) before starting a new one (file scanner.go, function main, line 111). 

* How do you determine if a goroutine has finished?

When a goroutine (which makes a call to the function checkConnection, file scanner.go, line 26) is ready to finish the checkConnection function reads from the channel c (file scanner.go, function checkConnection, line 38), allowing another goroutine to start. If the checkConnection function does not manage to finish by reading from the channel c, it has a timeout of 500 milliseconds (file scanner.go, function checkConnection, line 29) that will also finish the goroutine and free up the channel allowing another goroutine to start.


* Why is your solution deadlock free?

In the file scanner.go, function main, line 110 we write an integer (1) to the channel c. This happens before we start our goroutine. Then in the file scanner.go, function checkConnection, line 38 we read from the same channel. Before the goroutines start, a loop in file scanner.go, function main, line 115 starts writing to the channel. Then the goroutines start writing to the channel as well, so while the program is running there are always somewhere between 1 and 1024 inputs in the channel and the program ends with the channel being full. With this we can make sure that there is always something to read from the channel and a deadlock is not created by waiting forever trying to read from an empty channel that will never on written to again. That is one of the reasons our solution is deadlock free. The other reason is that in file scanner.go, function checkConnection, line 29 we have a Timeout for the function so that every goroutine ends after 500 milliseconds. If the function checkConnection (file scanner.go, line 26) does not finish and read from the channel c (file scanner.go, function checkConnection, line 38) it will not be able to run forever and block other goroutines from starting. What could happen if we did not have a timeout is that the channel could fill up and no one could write in it because there are routines that have written to the channel but never read from it.
    
* Do you manage to run as many goroutines as possible at any time?

Yes. In file scanner.go, function main, line 105 we set the limit of the channel c to 1024. That is the number of file descriptors that are free in our computers. With that limit the double loop in file scanner.go, function main, line 108 and 109 can write an integer (1) to the channel 1024 before waiting for the channel to be read from (a goroutine to finish) and the channel therefore run 1024 goroutines at a time (as many as possible).