# Dictator
#### Dictionary Generator is a dictionary generation tool which allows you to create dictionary lists in a very customized way.

![image](https://user-images.githubusercontent.com/70066388/109635222-6ce9f380-7b4a-11eb-94c8-c89fd2d9db78.png)

![image](https://user-images.githubusercontent.com/70066388/109636276-abcc7900-7b4b-11eb-90a2-983c8c92fe26.png)


Dictator is argument order sensitive, meaning that the output of one command will server as the input for the next. 
A few examples:

Take every line of list.txt, use different casings on each line, then for each different casing add '!', '@', '#', '%':

`Dictator --file ./list.txt --case --chars "!@#%"`



Take every line of list.txt, append either '_', '-' or nothing, then for each line generated add every line inside list.txt:

`Dictator --file ./list.txt --chars '_-' --file ./list.txt`


Generate the numbers 0-9, append each line of list.txt to each number, append the numbers 0-9 to each line, append special characters to each line twice:

`Dictator --nums nano --file ./list.txt --nums nano --chars '!@#%' --chars '!@#%'`

