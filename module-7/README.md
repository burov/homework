# Task write plain old Ping & Pong game

* Implement function `play` that accept array of two channels and returns the winner
* Every player should be runned in separate go routine
* Function `ballMissed` should be called each time where player get ball to check is this ball missed or not 
* If `ballMissed` player should count score to his oponent
* Function `gameEnded` can be used to check is game over or not
* Package `sync` might be very benefitial for syncronization primitives 
