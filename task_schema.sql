Drop table task;
Drop table users;
CREATE TABLE users ( 
ID BIGINT NOT NULL auto_increment,
Username char(32) UNIQUE NOT NULL,
Passwords  VARCHAR(100) NOT NULL,
primary key(ID)
);
CREATE TABLE task (
	ID BIGINT NOT NULL AUTO_INCREMENT,
	AssigneeID bigint NOT NULL ,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	Task VARCHAR(4096) NOT NULL,
	Note VARCHAR(1024),
    Done BOOL DEFAULT FALSE,
	PRIMARY KEY (ID),
    foreign key (AssigneeID) references users(ID)
);


