create table if not exists app_user
(
    id           bigserial primary key,
    account_name varchar(255) unique                                not null, -- 'account name'
    password     varchar(255)                                       not null, -- 'password'
    created_at   timestamp with time zone default CURRENT_TIMESTAMP not null, -- 'create time'
    updated_at   timestamp with time zone default CURRENT_TIMESTAMP not null, -- 'update time'
    is_deleted   boolean                  default false             not null  -- 'is deleted',
);
