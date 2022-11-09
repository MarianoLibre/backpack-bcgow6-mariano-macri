CREATE TABLE
  IF NOT EXISTS Library (
    Id int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    Address varchar(255) NOT NULL,
    PhoneNumber varchar(255),
    PRIMARY KEY (Id)
  );

CREATE TABLE
  IF NOT EXISTS Book (
    Id int NOT NULL AUTO_INCREMENT,
    Title varchar(255) NOT NULL,
    Quantity int NOT NULL,
    PRIMARY KEY (Id)
  );

CREATE TABLE
  IF NOT EXISTS User (
    Id int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    Age int,
    PRIMARY KEY (Id)
  );

CREATE TABLE
  IF NOT EXISTS Loan (
    Id int NOT NULL AUTO_INCREMENT,
    BookId int,
    UserId int,
    PRIMARY KEY (Id),
    FOREIGN KEY (BookId) REFERENCES Book (Id),
    FOREIGN KEY (UserId) REFERENCES User (Id)
  );

CREATE TABLE
  IF NOT EXISTS LibraryBook (
    Id int NOT NULL AUTO_INCREMENT,
    BookId int NOT NULL,
    LibraryId int NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (LibraryId) REFERENCES Library(Id),
    FOREIGN KEY (BookId) REFERENCES Book(Id)
  );

CREATE TABLE
  IF NOT EXISTS LibraryUser (
    Id int NOT NULL AUTO_INCREMENT,
    LibraryId int,
    UserId int,
    PRIMARY KEY (Id),
    FOREIGN KEY (LibraryId) REFERENCES Library (Id),
    FOREIGN KEY (UserId) REFERENCES User (Id)
  );
