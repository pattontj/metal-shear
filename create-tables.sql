DROP TABLE IF EXISTS vtuber;
CREATE TABLE vtuber (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    affiliation VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS clips;
CREATE TABLE clips (
    id INT AUTO_INCREMENT NOT NULL,
    link VARCHAR(128) NOT NULL,

    beginTime INT,
    endTime INT,

    vtuberID INT,

    PRIMARY KEY (`id`),
    FOREIGN KEY (VtuberID) REFERENCES vtuber(id)
);

INSERT INTO vtuber
(title, affiliation)
VALUES
('Inugami Korone', "Hololive"),
('Pomu Rainpuff', 'Nijisanji'),
('Kson', 'Indie'),
('Nina Kosaka', 'Nijisanji')