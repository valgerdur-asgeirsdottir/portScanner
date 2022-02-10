- How do you determine the number of goroutines scanning hosts concurrently?


- How and under what conditions do you start the goroutines?


- How do you determine if a goroutine has finished?


- Why is your solution deadlock free?
    In the file scanner.go, function main, line 110 we write 1 to the channel c. That happens before we start our goroutine. Then in the file scanner.go, function checkConnection, line 38 we read from the channel. While the goroutines are running we make sure to write enough into our channel, so we always have somewhere between 1 and 1024 inputs in the channel, that you can see in the file scanner.go, function main, line 115. With this we can make sure that there is always something to read from the channel and never anyone waiting forever to get an input that will never come. That is one of the reasons our solution is deadlock free. The other reason is that in file scanner.go, function checkConnection, line 29 we have a Timeout for the function so that every goroutine ends after some amount of time and can therefore not be going on forever and blocking other goroutines from happening. What could happen if we did not have a timeout is that the channel could fill up and no one could write in it because there is a routine that never reads out of it.
    
- Do you manage to run as many goroutines as possible at any time?
    Yes. In line ... we set the limit of the channel c to 1024. That is the number of file descriptors that are free in our computers. With that limit the for loop in line .. can add 1024 1s to the chanel and the chanel therefor run  1024 goroutines at a time.