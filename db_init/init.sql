CREATE SCHEMA IF NOT EXISTS jlpt
   CREATE TABLE IF NOT EXISTS words 
      (
         id SERIAL PRIMARY KEY,
         foreign1 VARCHAR(255) NOT NULL, 
         foreign2 VARCHAR(255) NOT NULL,
         definitions TEXT NOT NULL,
         level  VARCHAR(255) NOT NULL
      );


