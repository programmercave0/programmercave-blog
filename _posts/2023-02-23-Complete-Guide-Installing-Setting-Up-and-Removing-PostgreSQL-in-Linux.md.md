---
layout: post
title: "Complete Guide: Installing, Setting Up, and Removing PostgreSQL in Linux"
subtitle: "PostgreSQL is an open-source relational database management system. It is known for its reliability, scalability, and extensibility, making it a popular choice for enterprise-level applications. In this tutorial, we will cover how to install and set up PostgreSQL in Linux, how to create a database and table, how to insert values into the table, how to log in to a user and database, and how to completely remove PostgreSQL including all related files."
author: "Programmercave"
header-img: ""
tags:  [Linux, Arch-Linux, OpenSuse, Tumbleweed, Ubuntu, HowTo]
date: 2023-02-23
---

# Introduction

PostgreSQL is an open-source relational database management system. It is known for its reliability, scalability, and extensibility, making it a popular choice for enterprise-level applications. In this tutorial, we will cover how to install and set up PostgreSQL in Linux, how to create a database and table, how to insert values into the table, how to log in to a user and database, and how to completely remove PostgreSQL including all related files.

# Prerequisites

Before we begin, make sure you have the following prerequisites:

- Linux installed on your system
- A terminal emulator with root privileges

## Step 1: Installing PostgreSQL

To install PostgreSQL in Linux, follow these steps:

**Note that while the instructions for creating a database, inserting values, and logging in to a user and database are applicable to PostgreSQL on any Linux distribution, the installation and deletion commands used in this tutorial are specific to Arch Linux.**

1. Open the terminal emulator.

2. Update the package list:

```bash
sudo pacman -Syu
```

3. Install PostgreSQL using the following command:

```bash
sudo pacman -S postgresql
```

This will install PostgreSQL on your system.

## Step 2: Setting up PostgreSQL


After installing PostgreSQL, you need to perform some initial setup before you can use it. Follow these steps:

1. Initialize the PostgreSQL data directory using the following command:

```bash
sudo -iu postgres initdb -D /var/lib/postgres/data
```

This command initializes the PostgreSQL data directory with default settings.

2. Start the PostgreSQL service:

```bash
sudo systemctl start postgresql.service
```

3. Enable the PostgreSQL service:

```bash
sudo systemctl enable postgresql.service
```

This command ensures that PostgreSQL starts automatically on system boot.

## Step 3: Creating a Database and Table

After setting up PostgreSQL, you can create a database and table. Follow these steps:

1. Log in to the PostgreSQL server as the `postgres` user:

```bash
sudo -iu postgres psql
```

2. Create a new database using the following command:

```sql
CREATE DATABASE mydb;
```

This command creates a new database named `mydb`.

3. Connect to the newly created database using the following command:

```sql
\c mydb
```

4. Create a new table in the database using the following command:

```sql
CREATE TABLE mytable (id serial PRIMARY KEY, name VARCHAR(50), age INTEGER);
```

This command creates a new table named `mytable` with three columns: `id`, `name`, and `age`.

## Step 4: Inserting Values into the Table

After creating the database and table, you can insert values into the table. Follow these steps:

1. Insert a new row into the table using the following command:

```sql
INSERT INTO mytable (name, age) VALUES ('John Doe', 30);
```

This command inserts a new row into the `mytable` table with `name` set to `John Doe` and `age` set to `30`.

2. View the contents of the table using the following command:

```sql
SELECT * FROM mytable;
```

This command displays the contents of the `mytable` table.

## Step 5: Logging in to a User and Database

After creating a database and table, you may want to log in to a specific user and database. Follow these steps:

1. Log in to the PostgreSQL server as the `postgres` user:

```bash
sudo -iu postgres psql
```

2. Create a new user using the following command:

```sql
CREATE USER myuser WITH PASSWORD 'mypassword';
```

This command creates a new user named `myuser` with the password `mypassword`.

3. Grant permissions to the new user using the following command:

```sql
GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
```

This command grants all privileges to the `mydb` database to the `myuser` user.

4. Exit the PostgreSQL prompt using the following command:

```sql
\q
```

5. Log in to the PostgreSQL server as the myuser user:

```sql
psql -U myuser -d mydb
```

This command logs in to the `mydb` database as the `myuser` user.

## Step 6: Removing PostgreSQL

If you want to remove PostgreSQL from your Linux system, you can do so using the following steps:

1. Stop the PostgreSQL service:

```bash
sudo systemctl stop postgresql.service
```

2. Remove the PostgreSQL package using the following command:

```bash
sudo pacman -Rs postgresql
```

3. Remove the PostgreSQL data directory and configuration files using the following command:

```bash
sudo rm -r /var/lib/postgres/data
sudo rm -rf /etc/postgresql/
```

These commands will remove the PostgreSQL data directory (`/var/lib/postgres/`) and configuration files (`/etc/postgresql/`).

# Conclusion

PostgreSQL is a powerful and reliable relational database management system that is widely used in enterprise-level applications. In this tutorial, we covered how to install and set up PostgreSQL in Linux, how to create a database and table, how to insert values into the table, how to log in to a user and database, and how to completely remove PostgreSQL including all related files.

## References:

[PostgreSQL-ArchLinux](https://wiki.archlinux.org/title/PostgreSQL) <br>
[postgresql.org][def]

[def]: https://www.postgresql.org/