CREATE TABLE IF NOT EXISTS transactions (
    id serial primary key,
    sender int not null,
    receiver int not null,
    amount int not null,
    transfer_date timestamp not null default now(),
    foreign key (sender) references users(id),
    foreign key (receiver) references users(id)
);