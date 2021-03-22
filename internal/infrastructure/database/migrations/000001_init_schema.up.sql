CREATE TABLE IF NOT EXISTS forums(
    ID serial PRIMARY KEY,
    userID varchar(255),
    title varchar(255),
    topic varchar(255),
    forum_text text
);

CREATE TABLE IF NOT EXISTS comments(
    ID serial PRIMARY KEY,
    forumID serial,
    commentText text,
    CONSTRAINT fk_forum FOREIGN KEY(forumID) REFERENCES Forums(ID)
);