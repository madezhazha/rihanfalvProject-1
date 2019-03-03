

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
	thesislength integer,
	lables text
);

Create table japanthesis(
	thesisid serial primary key,
	thesistitle text,
	thesisauthor text,
	publicationtime date,
	thesiscontent text,
	thesislength integer,
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


--讨论区
CREATE TABLE topics (
  topicid serial primary key NOT NULL,
  posterid int2 NOT NULL,
  topictitle text NOT NULL,
  topiccontent text NOT NULL,
  creationtime timestamp(0) NOT NULL,
  numberofreplies int4 DEFAULT 0,
  finalreplytime timestamp(0),
  collectionvolume int4 DEFAULT 0,
  visitvolume int4 DEFAULT 0,
  japanorkorea int2,
  label varchar(255)
);


CREATE TABLE replies (
  replieid serial primary key NOT NULL,
  userid int4 NOT NULL,
  topicid int4 NOT NULL,
  replycontent text ,
  floor int4 NOT NULL,
  replytime timestamp(6)
);




