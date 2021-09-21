create table infoes(
	info_id serial not null primary key,
	from_where varchar(64) not null,
	to_where varchar(64) not null,	
	fly_time timestamp with time zone default current_timestamp,
	lnading_time timestamp with time zone default null,
	durring_fly date
);

insert into info(from_where, to_where,lnading_time)values(
'Buxoro', 'Tashkent', '2021-09-14 12:00'::timestamp);

insert into info(from_where, to_where,lnading_time)values(
'Uzbekistan', 'Russia', '2021-09-14 19:00'::timestamp);

insert into info(from_where, to_where,lnading_time)values(
'Khorazm', 'Buxoro', '2021-09-21 19:00'::timestamp);