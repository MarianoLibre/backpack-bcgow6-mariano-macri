CREATE TABLE IF NOT EXISTS library (
  LibraryId int NOT NULL AUTO_INCREMENT,
  Name varchar(255) NOT NULL,
  Address varchar(255) NOT NULL,
  PhoneNumber varchar(255),
  PRIMARY KEY(LibraryId)
);

CREATE TABLE IF NOT EXISTS Book (
  BookId int NOT NULL AUTO_INCREMENT,
  Title varchar(255) NOT NULL,
  Quantity int NOT NULL,
  PRIMARY KEY(BookId)
);

CREATE TABLE IF NOT EXISTS User (
  UserId int NOT NULL AUTO_INCREMENT,
  Name varchar(255) NOT NULL,
  Age int,
  PRIMARY KEY(UserId)
);

CREATE TABLE IF NOT EXISTS Loan (
  LoanId int NOT NULL AUTO_INCREMENT,
  BookId int,
  UserId int,
  PRIMARY KEY(LoanId),
  FOREIGN KEY(BookId) REFERENCES Book(BookId),
  FOREIGN KEY(UserId) REFERENCES User(UserId)
);

CREATE TABLE IF NOT EXISTS LibraryBook (
  LibraryBookId int NOT NULL AUTO_INCREMENT,
  LibraryId int,
  BookId int,
  FOREIGN KEY(LibraryId) REFERENCES Library(LibraryId),
  FOREIGN KEY(BookId) REFERENCES Book(BookId)
);

CREATE TABLE IF NOT EXISTS LibraryUser (
  LibraryUserId int NOT NULL AUTO_INCREMENT,
  LibraryId int,
  BookId int,
  FOREIGN KEY(LibraryId) REFERENCES Library(LibraryId),
  FOREIGN KEY(UserId) REFERENCES User(UserId)
);
