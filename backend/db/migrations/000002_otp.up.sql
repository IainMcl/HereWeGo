create table otp (
  id serial primary key,
  user_id int not null,
  code varchar(50) not null,
  created_at timestamp not null default now(),
  expires_at timestamp not null,

  constraint fk_otp_user_id foreign key (user_id) references users(id)
);
