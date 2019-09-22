/*stu测试数据*/
create table stu
(
	stuNo char(6) primary key,
	stuName varchar(10) not null,
	stuSex char(2) not null,
	stuAge tinyint not null ,
	stuSeat tinyint not null,
	stuAddress varchar(10) not null,
	ch tinyint,
	math tinyint 
)charset=utf8;

insert into stu values ('s25301','范建','男',18,1,'北京',80,null);
insert into stu values ('s25302','罗况','男',31,3,'上海',77,76);
insert into stu values ('s25303','申晶冰','女',22,2,'北京',55,82);
insert into stu values ('s25304','杜子腾','男',28,4,'天津',null,74);
insert into stu values ('s25305','史泰香','女',17,7,'河南',72,56);
insert into stu values ('s25318','郭迪辉','男',26,6,'天津',86,92);
insert into stu values ('s25319','拎壶冲','女',23,5,'河北',74,67);

insert into stu values ('s25320','Tom','男',24,8,'北京',65,67);
insert into stu values ('s25321','Tabm','女',23,9,'河北',88,77);

/*stuinfo测试数据*/
create table stuinfo
(
	stuNo char(6) primary key,
	stuName varchar(10) not null,
	stuSex char(2) not null,
	stuAge tinyint not null ,
	stuSeat tinyint not null,
	stuAddress varchar(10) not null
)charset=utf8;


insert into stuinfo values ('s25301','范建','男',18,1,'北京');
insert into stuinfo values ('s25302','罗况','男',31,3,'上海');
insert into stuinfo values ('s25303','申晶冰','女',22,2,'北京');
insert into stuinfo values ('s25304','杜子腾','男',28,4,'天津');
insert into stuinfo values ('s25305','史泰香','女',17,7,'河南');
insert into stuinfo values ('s25318','郭迪辉','男',26,6,'天津');
insert into stuinfo values ('s25319','拎壶冲','女',23,5,'河北');

/*stumarks测试数据*/

create table stumarks
(
examNo char(7) primary key,
stuNo char(6) not null ,
writtenExam int,
labExam int
);

insert into stumarks values ('s271811','s25303',80,58);
insert into stumarks values ('s271813','s25302',50,90);
insert into stumarks values ('s271815','s25304',65,50);
insert into stumarks values ('s271816','s25301',77,82);
insert into stumarks values ('s271819','s25318',56,48);
insert into stumarks values ('s271820','s25320',66,77);
