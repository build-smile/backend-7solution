db = db.getSiblingDB('my_app_db'); // ชื่อ DB ต้องตรงกับ config

db.createCollection('user'); // สร้าง collection

// (Optional) ใส่ document ตัวอย่าง
db.user.insertOne({
    name: "Admin",
    email: "admin@example.com",
    password: "hashed_password_here",
    createdAt: new Date()
});
