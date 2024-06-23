---
layout: post
title: "Programming paradigms - Procedural vs Object Oriented Programming"
description: "Notes on Programming paradigms from Scaler"
author: "Programmercave"
header-img: "/assets/images/LLD/programming-paradigms.jpg"
tags:  [System-Design, Software-Engineering]
date: 2024-06-21
toc: true
---
![Programming paradigms]({{ site.url }}/assets/images/LLD/programming-paradigms.jpg){:class="img-responsive"}

## Types of Programming paradigms

**Imperative:**
- an imperative program consists of commands for the computer to perform to change state e.g. C, Java, Python, etc.

**Declarative:**
- focuses on what the program should accomplish without specifying all the details of how the program should achieve the result e.g. SQL, Lisp, etc.

## Imperative

### Procedural programming:
- programming paradigm that uses a sequence of steps to solve a problem.
- based on the concept of the procedure call. 
- Procedures (a type of routine or subroutine) simply contain a series of computational steps to be carried out. 
- Any given procedure might be called at any point during a program's execution, including by other procedures or itself.
- **Think of all programming as managing the relationship between two fundamental concepts: state and behavior. State is the data of your program. Behavior is the logic.**
- State is held in data structures. Behavior is held in functions (also known as procedures or subroutines). A procedural application therefore passes data structures into functions to produce some output.
- Eg: Imagine you want to transfer some money from one account to another. These are the following steps that you would take:
  - Open the source account
  - Withdraw the money
  - Open the destination account
  - Deposit the money in destination account
  - A procedural version of this program would be:

#### Python Procedural Code

```python
 def transfer(source: int, destination: int, amount: int) -> None:
    source_account = get_account(source)
    update_account(source_account, -amount)
    destination_account = get_account(destination)
    update_account(destination_account, amount)

def get_account(number: int) -> dict:
    return list(filter(lambda account: account['number'] == number,accounts))[0]

def update_account(account: int, delta: int) -> None:
    account['balance'] += delta
```

### Object Oriented Programming

- programming paradigm that uses objects to model real-world things and aims to implement state and behavior using objects.
- State and behaviour are combined into one new concept: an Object. An OO application can therefore produce some output by calling an Object, without needing to pass data structures.
- Advantages of OO include the potential for information hiding: if a caller needn't pass any data structure, then the caller needn't be aware of any data structure, and can therefore be completely decoupled from the data format.

#### Abstraction:
- process of hiding the implementation details of a program from the user.

###### Advantages of Abstraction
- used to create a boundary between the application and the client code.
- separate responsibilities into software entities (classes, method, etc.) that only know the required functionality of each other but not how that functionality is implemented.
- It allows the programmer to change the internal implementation of methods or concrete classes without hampering the interface.
-  increase the code security as only relevant details will be provided to users.

#### Encapsulation:
- used to hide the values or state of a structured data object inside a class,
- preventing direct access to them by clients in a way that could expose hidden implementation details or violate state invariance maintained by the methods.

###### Advantages of Encapsulation:
- **Hiding Data** - Users will have no idea how classes are being implemented or stored. All that users will know is that values are being passed and initialized.
- **More Flexibility** - Enables you to set variables as read or write-only.
- **Easy to Reuse** - With encapsulation it's easy to change and adapt to new requirements.

#### Class:
- is a blueprint which you use to create objects.

#### Object:
- is an instance of a class.

#### Java OOP Code
```java
public class OopBankAccount {
    private Integer number;
    private Integer balance;
    
    public OopBankAccount(Integer number, Integer balance) {
        this.number = number;
        this.balance = balance;
    }
    
    void deposit(Integer amount) {
        this.balance += amount;
    }
    
    void withdraw(Integer amount) {
        this.balance += amount;
    }
    
    void transfer(OopBankAccount destination, Integer amount) {
        this.withdraw(amount);
        destination.deposit(amount);
    }
}
```

#### Python OOP Code
```python
class OopBankAccount:
	def __init__(self, balance, number):
		self.__number = number
		self.__balance = balance

	def getNumber(self):
		return self.__number

	def setNumber(self, number):
		self.__number = number 

	def getBalance(self):
		return self.__balance

	def setBalance(self, balance):
		self.__balance = balance

	def deposit(self, amount):
		self.__balance += amount

	def withdraw(self, amount):
		self.__balance -= amount

	def transfer(self, destination, amount):
		self.withdraw(amount)
		destination.deposit(amount)
```

#### Golang Code
```go
type BankAccount struct {
	AccountNumber int64
	Name          string
	Balance       float64
}

type BankAccountOps interface {
	GetAccountNumber() (int64, error)
	GetName() (string, error)
	GetBalance() (float64, error)
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Transfer(destiantion *BankAccount, amount float64) error
	PrintBalance()
}

func NewBankAccount(number int64, name string, balance float64) *BankAccount {
	return &BankAccount{AccountNumber: number, Name: name, Balance: balance}
}

func (b *BankAccount) GetAccountNumber() (int64, error) {
	return b.AccountNumber, nil
}

func (b *BankAccount) GetName() (string, error) {
	return b.Name, nil
}

func (b *BankAccount) GetBalance() (float64, error) {
	return b.Balance, nil
}

func (b *BankAccount) Deposit(amount float64) error {
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	b.Balance -= amount
	return nil
}

func (b *BankAccount) Transfer(destiantion *BankAccount, amount float64) error {
	b.Withdraw(amount)
	destiantion.Deposit(amount)

	return nil
}

func (b *BankAccount) PrintBalance() {
	fmt.Println(b.AccountNumber, b.Name, b.Balance)
}
```

#### Advantages:
- **Reusability**: Through classes and objects, and inheritance of common attributes and functions.
- **Security**: Hiding and protecting information through encapsulation.
- **Maintenance**: Easy to make changes without affecting existing objects much.
- **Inheritance**: Easy to import required functionality from libraries and customize them, thanks to inheritance.

#### Disadvantages:
 - Beforehand planning of entities that should be modeled as classes.
 - OOPS programs are usually larger than those of other paradigms.
- [Banana-gorilla problem](https://dev.to/efpage/what-s-wrong-with-the-gorilla-2l4j#:~:text=Joe%20Armstrong%2C%20the%20principal%20inventor,and%20the%20entire%20jungle.%22.) - You wanted a banana but what you got was a gorilla holding the banana and the entire jungle

> These are my notes from the Low-Level Design (LLD) course I took at Scaler.

> Check Python, Java and Go code on [Github Repo](https://github.com/abhiabhi0/scaler-fundamentals/tree/master/oop/code)





