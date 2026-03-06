CREATE TABLE notifications (

    id UUID PRIMARY KEY,
    user_id TEXT,
    type TEXT,
    message TEXT,
    status TEXT,
    created_at TIMESTAMP
);