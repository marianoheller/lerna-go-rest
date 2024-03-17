CREATE TABLE articles (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title  TEXT NOT NULL,
    slug  TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
INSERT INTO articles (id, user_id, title, slug)
VALUES ('1', 100, 'Hi', 'hi'),
    ('2', 200, 'sup', 'sup'),
    ('3', 300, 'alo', 'alo'),
    ('4', 400, 'bonjour', 'bonjour'),
    ('5', 500, 'whats up', 'whats-up');