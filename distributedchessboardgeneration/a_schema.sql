CREATE TABLE chessboards (
    id INTEGER PRIMARY KEY,
    chessboard TEXT NOT NULL
);

CREATE TABLE move (
    id INTEGER PRIMARY KEY,
    from_id INTEGER NOT NULL REFERENCES chessboards(id)
        ON DELETE CASCADE,
    -- Note, the foreign key has been turned off due to conflicts between assignments
    to_id INTEGER NOT NULL
);