create table orders(
	orderNumber serial,
	itemId int,
	price int,
	customerId int,
	vehicleId int,
	isDelivered boolean default FALSE,
	primary key(orderNumber),
	CONSTRAINT fk_itemId
      FOREIGN KEY(itemId)
	  REFERENCES items(itemId),
	CONSTRAINT fk_customerId
      FOREIGN KEY(customerId)
	  REFERENCES customers(customerId),
	CONSTRAINT fk_vehicleId
      FOREIGN KEY(vehicleId)
	  REFERENCES vehicles(vehicleId)
);