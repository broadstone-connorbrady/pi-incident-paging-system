create table users
(
    id         serial
        constraint users_pk
            primary key,
    uuid       uuid      default gen_random_uuid(),
    email      varchar(36) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);

create unique index users_email_uindex
    on users (email);

create unique index users_uuid_uindex
    on users (uuid);

create table pagers
(
    id         serial
        constraint pagers_pk
            primary key,
    uuid       uuid      default gen_random_uuid(),
    address    varchar(12) not null,
    user_id    int         not null
        constraint pagers_users_id_fk
            references users,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);

create unique index pagers_address_uindex
    on pagers (address);

create unique index pagers_uuid_uindex
    on pagers (uuid);