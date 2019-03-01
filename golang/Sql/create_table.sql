

--homepage
create table homepagenews(
imgurl varchar(200),
linkurl varchar(200),
title varchar(200),
id serial);

create table homepageartical (
id serial,
imgurl varchar(200),
linkurl varchar(200),
brif    varchar(200),
date date);

--head
create table Users(
	userid serial primary key,
	username varchar(30),
	password varchar(20),
	email varchar(30),
	image varchar(100),
	integal  integer,
	registrationdate varchar(100)
);

--论文
Create table koreathesis(
	thesisid serial primary key,
	thesistitle text,
	thesisauthor text,
	publicationtime date,
	thesiscontent text,
	thesislength text,
	lables text
);

Create table japanthesis(
	thesisid serial primary key,
	thesistitle text,
	thesisauthor text,
	publicationtime date,
	thesiscontent text,
	thesislength text,
	lables text
);

--收藏
create table collection(
   userid  serial primary key,
   collectionid bigint,
   collectiontype text,
   collectioncontentid  bigint,
   collectiontime  date
);


--feedback
create table feedback (
feedbackId  serial PRIMARY KEY,
userId bigint NOT NULL,
feedbackContent text NOT NULL,
feedbackType varchar(15) NOT NULL ,
feedbacktime timestamp with time zone default current_timestamp ,
feedbackreplie text NOT NULL
);



--案例分析



