use douyin;

create table messages (
    id int auto_increment,
    from_user_id int not null,
    to_user_id int not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index idx_from_user_id (from_user_id),
    index idx_to_user_id (to_user_id),
    primary key(id)
);