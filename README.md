# Hackathon-Practice
My university is trying to start a hackathon team, and I will be putting all of my practice problems under this repository.



# Dutch National Flag problem (DNFprob.go)
Given an array containing only 0’s, 1’s, and 2’s, sort it in linear time and using constant space.

The trick was all in the 3 for loops I used in the function to build the array.
This method could be used to solve any sorting problem where the elements in the array are known. 


# Tidy Numbers (Tidy_nums.go)
Tidy Numbers, from Google Code Jam 2017
Tatiana likes to keep things tidy. Her toys are sorted from smallest to largest, her pencils are
sorted from shortest to longest and her computers from oldest to newest. One day, when
practicing her counting skills, she noticed that some integers, when written in base 10 with no
leading zeroes, have their digits sorted in non-decreasing order. Some examples of this are 8,
123, 555, and 224488. She decided to call these numbers tidy. Numbers that do not have this
property, like 20, 321, 495 and 999990, are not tidy.
She just finished counting all positive integers in ascending order from 1 to N. What was the last
tidy number she counted?

This was listed as one of the "hard" type problems for the practice. 
A bunch of different methods were listed as well. 
I was not satisfied with any of these, so I applied a lot of the skills I have learned as a math major. Over the course of a week I came up with my own method to solve the problem.
Through lots of thought, I realized that a large amount of cases could be reduced to simpler ones by knowing which digit caused the number to become "un-tidy".

For example --- 123457000 => 7000 or 1234567890 => 90 

This works since there is a tidy number every 10^k numbers.
The only special case is when there are repeating digits right before the number becomes "un-tidy" EX 12345555500 or 12222221,
But this case can once again be reduced to the first digit where it started repeating EX 12345777773 => 12345700000 ( The proof of this will be more complex but still relies on the fact that there exists a tidy number every 10^k numbers).

As of now, I have yet to formally prove my method (I probably will in the future in a LaTeX document) but it appears to work well. 





