DROP TABLE IF EXISTS vtuber;
CREATE TABLE vtuber (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    channel VARCHAR(256) NOT NULL,
    affiliation VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS clips;
CREATE TABLE clips (
    id INT AUTO_INCREMENT NOT NULL,
    link VARCHAR(256) NOT NULL,

    beginTime INT,
    endTime INT,

    vtuberID INT,

    PRIMARY KEY (`id`),
    FOREIGN KEY (VtuberID) REFERENCES vtuber(id)
);

INSERT INTO vtuber
(title, channel, affiliation)
VALUES
('inugami korone', 'https://www.youtube.com/channel/UChAnqc_AY5_I3Px5dig3X1Q', "hololive"),
('pomu rainpuff', 'https://www.youtube.com/channel/UCP4nMSTdwU1KqYWu3UH5DHQ', 'nijisanji'),
('kson', 'https://www.youtube.com/c/ksonONAIR', 'Indie'),
('nina kosaka', 'https://www.youtube.com/channel/UCkieJGn3pgJikVW8gmMXE2w', 'nijisanji');


INSERT INTO clips
(link, beginTime, endTime, vtuberID)
VALUES
('https://www.youtube.com/watch?v=5VWWIXD4mHE', '0', '0', 
    (SELECT id FROM vtuber WHERE title ='nina kosaka')
)