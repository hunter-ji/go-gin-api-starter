create table if not exists app_user
(
    id           bigserial primary key,
    account_name varchar(255) unique                                not null, -- 'account name'
    password     varchar(255)                                       not null, -- 'password'
    created_at   timestamp with time zone default CURRENT_TIMESTAMP not null, -- 'create time'
    updated_at   timestamp with time zone default CURRENT_TIMESTAMP not null, -- 'update time'
    is_deleted   smallint                 default 0                 not null  -- 'is deleted',
);

insert into app_user (account_name, password)
values ('hunter', '69a329523ce1ec88bf63061863d9cb14');