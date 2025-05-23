-- สร้างตาราง users (IT Support/Admin)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    name VARCHAR(128),
    email VARCHAR(128),
    role VARCHAR(32) DEFAULT 'it_support',
    created_at TIMESTAMP DEFAULT NOW()
);

-- สร้างตาราง tickets (ข้อมูล ticket ที่ user แจ้ง)
CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(32) DEFAULT 'open', -- open, in_progress, resolved, closed
    created_by_ip VARCHAR(64),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by INT,
    FOREIGN KEY (updated_by) REFERENCES users(id)
);

-- สร้างตาราง ticket_logs (ประวัติการแก้ไข/คอมเมนต์)
CREATE TABLE ticket_logs (
    id SERIAL PRIMARY KEY,
    ticket_id INT NOT NULL,
    user_id INT, -- nullable: อาจเกิดจากระบบ
    action VARCHAR(64),
    content TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (ticket_id) REFERENCES tickets(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- สร้างตาราง attachments (ไฟล์แนบ)
CREATE TABLE attachments (
    id SERIAL PRIMARY KEY,
    ticket_id INT NOT NULL,
    file_url VARCHAR(255),
    uploaded_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (ticket_id) REFERENCES tickets(id)
);

-- สร้างตาราง notifications (บันทึกการแจ้งเตือน)
CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    ticket_id INT,
    sent_to VARCHAR(128),
    message TEXT,
    status VARCHAR(32),
    sent_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (ticket_id) REFERENCES tickets(id)
);
