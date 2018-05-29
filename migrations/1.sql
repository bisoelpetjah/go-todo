use todo;

create table todo (
    id int not null auto_increment,
    message text,
    is_done tinyint(1) default 0,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    primary key (id)
);
