create domain alphanum as varchar(20) check (value ~ '^[A-Z0-9]+$');
CREATE TYPE valid_type AS ENUM ('bike', 'truck');

create table vehicles(
	vehicleId serial primary key,
	registrationNumber alphanum unique,
	vehicleType valid_type,
	city text,
	activeOrdersCount int check (activeOrdersCount<3) default 0
);