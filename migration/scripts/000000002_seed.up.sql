-- Dữ liệu demo
INSERT INTO profiles (id, name, email, department) VALUES
    ('alice', 'Alice Nguyen', 'alice@example.com', 'Finance Ops'),
    ('bob',   'Bob Tran',     'bob@example.com',   'Treasury'),
    ('carol', 'Carol Le',     'carol@example.com', 'Support')
ON CONFLICT (id) DO NOTHING;
