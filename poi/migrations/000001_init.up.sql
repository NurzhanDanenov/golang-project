create table users (
    id       serial       not null unique,
    name     varchar(255) not null,
    email    varchar(255) not null unique,
    age      int          not null,
    gender   varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE images
(
    id          serial       not null unique,
    image       varchar(255) not null
);

CREATE TABLE users_images
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    image_id int references images (id) on delete cascade not null
);