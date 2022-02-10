- How do you determine the number of goroutines scanning hosts concurrently?


- How and under what conditions do you start the goroutines?


- How do you determine if a goroutine has finished?
    

- Why is your solution deadlock free?

    
- Do you manage to run as many goroutines as possible at any time?
    Yes. In line ... we set the limit of the channel c to 1024. That is the number of file descriptors that are free in our computers. With that limit the for loop in line .. can add 1024 1s to the chanel and the chanel therefor run  1024 goroutines at a time.