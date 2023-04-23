CREATE TABLE agents
(
    fio varchar(255) not null,
    login varchar(255) not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE shops
(
    id varchar(255) not null unique,
    title varchar(255) not null
);

CREATE TABLE market_places
(
    id varchar(255) not null unique,
    title varchar(255) not null,
    shop_id varchar(255) references shops (id) on delete cascade not null
);

CREATE TABLE credits
(
    id serial not null unique,
    title varchar(255) not null,
    summary varchar(255) not null,
    timelimit varchar(255) not null,
    agent_id varchar(255) references agents (login)  on delete cascade not null,
    m_place_id varchar(255) references market_places (id) on delete cascade not null
);

CREATE TABLE agents_market_places
(
    id serial not null unique,
    agent_id varchar(255) references agents (login)  on delete cascade not null,
    m_place_id varchar(255) references market_places (id) on delete cascade not null
);