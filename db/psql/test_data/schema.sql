CREATE TABLE cities (
    city_id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    city_name text,
    city_alias varchar(3),
    city_state text
);

CREATE TABLE teams (
    team_id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    city_id INT,
    team_name text,
    PRIMARY KEY(team_id),
    CONSTRAINT fk_city
        FOREIGN KEY(city_id)
        REFERENCES cities(city_id)
);

CREATE TABLE players (
    player_id INT UNIQUE GENERATED ALWAYS AS IDENTITY,
    team_id INT,
    first_name text,
    last_name text,
    jersey_number varchar(2),
    PRIMARY KEY(player_id),
    CONSTRAINT fk_team
        FOREIGN KEY(team_id)
        REFERENCES teams(team_id)
);