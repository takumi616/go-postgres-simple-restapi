CREATE DATABASE english_vocabulary;

create table vocabulary(
   id       serial     PRIMARY KEY,
   title    varchar(30)   not null,
   sentence varchar(200)  not null,
   meaning  varchar(100)  not null
);


 