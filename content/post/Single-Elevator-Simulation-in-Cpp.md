---
author: Programmercave
date: "2020-04-16T00:00:00Z"
description: Here we are going to implement the Single Elevator Simulation in C++.
  Initially the elevator is at ground floor. It is represented by 0. Floors below
  ground floor are represented by negative integers.The elevator has maximum capacity
  it can carry, maximum and minimum floor it can carry to.The elevator accepts the
  request of floor the passengers want to go. If the elevator is empty then the first
  request sets the direction of the elevator. After that it checks whether requests
  are valid. For eg. if the direction of elevator is UP, current floor is 3 and if
  the passenger enters request for floor -1 then the request is discarded and the
  passenger does not enter in the elevator.
header-img: /assets/images/2020-04-16-Single-Elevator-Simulation-in-Cpp/c++_elevator_simulation.jpg
tags:
- Cpp
- Project
title: Single Elevator Simulation in C++
---



Here we are going to implement the **Single Elevator Simulation in C++**. Initially the elevator is at ground floor. It is represented by 0. Floors below ground floor are represented by negative integers.
The elevator has maximum capacity it can carry, maximum and minimum floor it can carry to.

The elevator accepts the request of floor the passengers want to go. If the elevator is empty then the first request sets the direction of the elevator. After that it checks whether requests are valid. For eg. if the direction of elevator is UP, current floor is 3 and if the passenger enters request for floor -1 then the request is discarded and the passenger does not enter in the elevator.

Below is the class diagram of Elevator class.

![C++ Elevator Simulation](/assets/images/2020-04-16-Single-Elevator-Simulation-in-Cpp/c++_elevator_simulation.jpg)

(`-`) means the attributes and methods are private. (`+`) means the attributes and methods are public.

`Direction` is the `enum` data type which can take the one of the two values (`UP`, `DOWN`). `direction` is the enum variable.

`requests` is an integer `vector` which store the requests entered by passengers. `min_floor` and `max_floor` store the minimum floor and maximum floor in the building respectively.

`current_floor` stores the floor at which the elevator is currently. Initially it is 0 that means the elevator is at ground floor.

`passengers` stores the current number of passengers in the elevator. Initially it is 0.

`capacity` stores the maximum number of people the elevator can carry.


```cpp
class Elevator
{
	enum Direction{ UP, DOWN };
	Direction direction;

	std::vector<int> requests = {};
 	int min_floor; //Undergroud floor will be shown using negative value
 	int max_floor;
 	int current_floor = 0; //ground floor
 	std::size_t passengers = 0;
 	std::size_t capacity;

 public:
 	Elevator(int min_floor, int max_floor, std::size_t capacity) :
 			min_floor(min_floor),
 			max_floor(max_floor),
 			capacity(capacity)
 			{}
 	~Elevator() {}

 	void start_elevator();

 private:
 	void set_request();
 	int check_request(int floor) const;
 	int is_valid_request(int floor);
 	void set_direction(int floor);
 };
```

The constructor accepts the minimum floor, maximum floor in the building and the capacity of the elevator. 

`check_request()` function checks whether the passenger’s request is valid. The floor passenger wants to visit is passed as an argument and if the request is valid it returns 0.

There are many reasons for an invalid request. If the elevator is going upward and passenger wants to go downward then 1 is returned. If the elevator is going downward and passenger wants to go upward then 2 is returned. If the floor entered by passenger does not exist then 3 is returned.

Here is C++ implementation of `check_request()` :

```cpp
int Elevator::check_request(int floor) const
{
 	if (passengers != 0 && direction == UP && floor < current_floor)
 	{
 		return 1;
 	}
 	else if (passengers != 0 && direction == DOWN && floor > current_floor)
 	{
 		return 2;
 	}
 	else if (floor > max_floor || floor < min_floor)
 	{
 		return 3;
 	}
 	else
 	{
 		return 0;
 	}
}
```

`check_request()` is a constant function because it does not modify any variable.

<br/>

`is_valid_request()` calls `check_request()` and outputs the error and returns the issue number.

```cpp
int Elevator::is_valid_request(int floor)
{
 	int issue_num = check_request(floor);

 	if (issue_num == 1)
 	{
 		std::cout << "Elevator is going UP.\n";
 	}
 	else if (issue_num == 2)
 	{
 		std::cout << "Elevator is going DOWN.\n";
 	}
 	else if (issue_num == 3)
 	{
 		std::cout << "This floor does not exist\n";
 	}
 	return issue_num;
}
```

When the elevator is empty, the first valid floor number sets the direction of the elevator. If the entered floor number is greater than the current floor number than the direction is upward else the direction is downward. This is done with the help of `set_direction()` function in which floor number is passed as an argument.

```cpp
void Elevator::set_direction(int floor)
{
 	if (floor > current_floor)
 	{
 		direction = UP;
 	}
 	else if (floor < current_floor)
 	{
 		direction = DOWN;
 	}
}
```

`set_request()` function accept passengers request, check their validity and then add them in requests array. String variable `dest_floors_str` store all requests for eg “1 2 3 4 7”. `dest_floor_str` store request for a single floor for eg. “1” or “2” or “3”. Integer variable `dest_floor` store request for a single floor as integer.

`stringstream` variable `sstream` helps to enter single floor request into `dest_floor_str` from `dest_floors_str`. When no request is entered, enter “GO” to go to next floor. Also to exit from the program when the elevator is empty enter “GO”.

`is_valid_request()` is called to check whether entered floor is valid. If it is valid then we check the number of passengers in the elevator. If the number of passenger is zero means this is the first request and it is used to set the direction of the elevator, so `set_direction()` is called. After than this request is added to the `requests` array and number of passenger is incremented.

If the number of passengers in the elevator is equal to the capacity of the elevator, means we cannot add any more request so we exit from this function.

<br/>

Here is the C++ implementation of `set_request()` :

```cpp
void Elevator::set_request() 
{
	std::string dest_floors_str; //stores all floors request
 	std::string dest_floor_str; //stores single floor in string
 	int dest_floor; // stores single floor as integer

 		
 	std::size_t num_of_reqs = capacity - passengers;
	std::cout << "\n" << num_of_reqs << " passengers can enter in the elevator right now\n";
	
	std::cout << "\nEnter \"GO\" if no one enters from the floor \nOr to exit from program if elevator is idle\n";

 	std::cout << "\nEnter destination floor number.\n";

 	std::getline(std::cin, dest_floors_str);
 	std::stringstream sstream(dest_floors_str);

 	while (sstream >> dest_floor_str)
 	{
 		if (dest_floor_str == "GO" || dest_floor_str == "Go" || dest_floor_str == "go" || dest_floor_str == "gO")
 		{
 			return;
 		}
 		else
 		{
 			dest_floor = std::stoi(dest_floor_str);
 			if (passengers < capacity)
 			{
 				int is_valid = is_valid_request(dest_floor);
 				if (is_valid == 0)
 				{
 					if (passengers == 0)
		 			{
 	 					set_direction(dest_floor);
 	 				}
 	 				requests.push_back(dest_floor);
 					passengers++;
				}
 			}
 			else if (passengers == capacity)
 			{
 				std::cout << "Elevator full!! Cannot accept more requests\n";
 				return;
 			}
 		}
 	}
}
```

`start_elevator()` is the function which is called in the `main()` function. It calls `set_request()` which enters valid requests in the requests array. Then the array is sorted. For eg. the requests are entered in [4, 2, 3, 1] after sorting they are [1, 2, 3, 4] because first the elevator will stop at 1 then 2 ans so on. 

A `while` loop is run until all requests are executed. Integer variable `next_floor` is set with the next floor to visit. Then all requests of the `next_floor` is deleted from the array and respective number of `passengers` are decremented. Then this `next_floor` becomes the `current_floor`.

If elevator is at the extreme floor then the direction is reversed. 

When the floor is reached at floor 1, again requests are entered in the array and the array stores [2, 3, 4, 6, 4, 7]. To let this happen and to run elevator continuously again `set_request()` is called in the while loop.

Here is the C++ implementation of `start_elevator()` :

```cpp
void Elevator::start_elevator()
{
 	std::cout << "\nFLOOR : " << current_floor << "\tNumber of Occupants : " << passengers <<"\n";

 	//Entering requests for first time
 	set_request(); 
 	std::sort(requests.begin(), requests.end());
 	int next_floor;

 	while (!requests.empty())
 	{
 		if (direction == UP)
 		{
 			next_floor = requests[0];
 		}
 		else if (direction == DOWN)
 		{
 			next_floor = requests[requests.size() - 1];
 		}

 		auto next_floor_req = std::find(requests.begin(), requests.end(), next_floor);
 		while (next_floor_req != requests.end())
 		{
 			requests.erase(next_floor_req); //removing next floor's requests
 			passengers--;
 			next_floor_req = std::find(requests.begin(), requests.end(), next_floor);
 		}
 		current_floor = next_floor;

 		std::string dir;
 		if (direction == UP)
 		{
 			dir = "UP";
 		}
 		else
 		{		
 			dir = "DOWN";
 		}

 		//Entering requests for current floor
 		std::cout << "\n=======================================================\n"
    		"FLOOR : " << current_floor 
    		<< "\tNumber of Occupants : " << passengers 
    		<< "\n\nDIRECTION : " << dir 
    		<< "\tTotal Capacity : " << capacity
    		<< "\n\nMinimum floor number is " << min_floor
    		<< "\tMaximum floor number is " << max_floor
    		<< "\n\n=======================================================\n";

 		if (current_floor == max_floor)
 		{
 			direction = DOWN;
 		}
		else if (current_floor == min_floor) 		
		{
 			direction = UP;
 		}

 		set_request();
 		std::sort(requests.begin(), requests.end());	
 	}
}
```

In `main()` function maximum floor, minimum floor and capacity of the elevator is entered. Then `start_elevator()` is called.

```cpp
int main()
{
	std::string capacity_str, min_floor_num_str, max_floor_num_str;
	int min_floor_num, max_floor_num;
	std::size_t capacity;

	std::cout << "Enter minimum floor number, maximum floor number in the building\n";
	std::cin >> min_floor_num_str;
	std::cin >> max_floor_num_str;

	min_floor_num = std::stoi(min_floor_num_str);
	max_floor_num = std::stoi(max_floor_num_str);

	std::cout << "Enter capacity for the elevator\n";
	std::cin >> capacity_str;
	std::cin.ignore();
	std::stringstream capacity_stream(capacity_str);
	capacity_stream >> capacity;

	Elevator elevator(min_floor_num, max_floor_num, capacity);
	elevator.start_elevator();	
}
```

Command to compile `g++ elevator_design.cpp -o elevator_design`

And to run `./elevator_design`

Get the full code on [Github](https://github.com/{{< param "github_username" >}}/Small-Projects/blob/master/Elevator-Design-C++-Implementation/elevator_design.cpp)

If you have any improvements or suggestions then please comment below.

<h3>Other posts worth visiting:</h3>

[C++: Tic Tac Toe](/C-Tic-Tac-Toe)<br/>

[C++: Simple Pendulum Animation on Ubuntu Machine](/C-Simple-Pendulum-Animation-on-Ubuntu-Machine)<br/>
