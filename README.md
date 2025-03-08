# HANGMAN

The game's goal is simple : find a word letter by letters, before the hangman drawing is completed.    
For our hangman, you have 10 attempts before it's too late and you lose !

## Table of contents 
1. [**Rules**](#-rules)
2. [**Features**](#-features)
3. [**How to run**](#-how-to-run)
4. [**Organisation**](#-organisation)
5. [**Tree structure**](#-tree-structure)
6. [**Technology**](#-technology)
7. [**Contributors**](#-contributors)
 --- 
## **RULES**  
The word you need to find will be hiddden behind " _ ", each represent a letter of the word.  
You can guess only a letter at a time, or the whole word, but it cost more!  
Each wrong guess add a part of the drawing.  
If you find the word before the drawing is complete, you win, otherwise you lose.

--- 

## **FEATURES**  
Random words to be chosen from a .txt file.  
Restart allowed.  
Enter a letter.

---
## **How to run**
You will need to have Go installed. 

Clone the repository :   
```bash
git clone https://github.com/xLuuxx/Hangman-Classic.git
cd Hangman-Classic
go run .
```
___

## **ORGANIZATION**  
*DIRECTORY* **data** : which contains our .txt files : words.txt / hangman.txt / standard.txt / save.txt.  
*DIRECTORY* **cmd** : which contains our main.go file.  
*DIRECTORY* **internal** : which contains our function in " .go " : features.go / input.go / game.go / display.go    
*DIRECTORY* **pkg** : which contains a directory utils with a utils.go file.

## **TREE STRUCTURE**
```
hangman-classic/  
├── cmd  
│    └── main.go  
├── internal  
│    ├── features.go  
│    ├── input.go  
│    ├── game.go  
│    ├── display.go  
│    ├── hangman.go
│    ├── random.go
└── data  
│    ├── words.txt   
│    ├──hangman.txt  
│    ├──save.txt  
└──  └──standard.txt
```
---

## **TECHNOLOGY**  
Goland, Go.
---
**CONTRIBUTORS**  
Elisabeth  
Luu

--- 
[TRELLO](https://trello.com/b/kz66rQZh/projet-hero)
---
**THANKS**  
Thank you for reading our readme.  
Thanks to the school's mentors and classmates for their help. 
