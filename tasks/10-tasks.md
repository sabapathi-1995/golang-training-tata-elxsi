- Take the file data.txt, 
- Using Generator pattern read the file
- But read the each line of the file and send to a channel 
- Dont use any readrs or bufio, you need to find the line by your ownlogic
- once a new line is found send it to the channel , each line should be given a line number


- 1. I want the same line to be send to multiple channels
- 2. One go routine read from the channel find out the number of words
- 3. One go routine read the line from the channel and find out number of chars
- 4. One go routine read the line from the channel and find out the overall (from all the channels), the number of words used
- 5. the fourth one can be multiple workers (should be able to increse the number of workds)





