# Go programming note make by dungad2k2
<div style="text-align: center;">
  <img src="golang.png" alt="Example Image" width="300" height="200">
</div>


## Chapter 1: Programming a Computer

Provide some basic knowledge about computer

### Computer and its components:

A Computer is consists of 4 main parts:
- The **memory unit** (MU): where data and programs are stored.
   - For instance, we can store into the memory unit the grades of a college class,....
- The **arithmetic and logic unit** (ALU): perform arithmetic and logical operations on data stored into the memory unit. This unit can perform for instance additions, incrementations, decrementations,...
   - Ex: When you need find sum of 2 integer, ALU will execute add operation.
- The **input and output unit (I/OU)** will be charge of loading data into the memory unit from an input device. This unit also sends data from the memory unit to an output device
- The **control unit** will receive instructions from programs and will control the activity of the other units. 

### Memory: 
A computer is composed of two types of memory: 
- The central memory
- The auxiliary memory
  
Two categories of memory exist:
- Volatile
- Non volatile

1. The central memory:
This type of memory is necessary to run the operating systems and all the other programs your computer will run. The central memory contains two types of storage: 
   - RAM (Random Access Memory): when you turn off computer the memory contained in this type of storage will be erased. The operating system and the programs you use will be loaded into this memory -> **volatile**
   - ROM (Read-only memory): contains data necessary for the computer to run correctly. This kind of memory is **not volatile** (when turn of compute, data will not be erased). It's designed to be only readable and not updated by the system.
2. The auxiliary memory: 
This type of memory is **not volatile**. When the power is going off, the data stored will not be earased. Some of auxiliary memory such as: CD-ROM, DVD,....Read and writes to this type of memory is **slow** compared to the RAM

### CPU
CPU is standed for **Central Processing Unit**. The CPU is also denoted **processor**. The CPU contains:
- Arithmetic and logic unit (ALU)
- Control unit (CU)

The CPU will be responsible for executing the instructions given by a program. For instance, the program can instruct to perform an addition between two numbers. Those numbers will be retrieved from the memory unit and passed to the ALU. The program might also require performing an I/O operation like reading data from the hard drive and loading it into the RAM for further processing. The CPU will execute those instructions.

### What is a program ?

To make computers do something, we have to feed them with precise instructions. This set of instructions is called "program". 

### How to speak to the machine ?

Intructions that are given to the machine are written with programming languages. Go is a programming language like: C/C++, Java, .....

There are two types of programming languages:
1. Low level
2. High level

Low level programming languages are closer to the processing unit's instructions. Higher-level languages provide constructs that make them easier to learn and to use. Some high-level languages are compiled, others are interpreted and some are in between. When source files are written, the program that they define cannot be executed immediately. The source file needs to be compiled by using a compiler. The compiled will transform source files into an executable. 