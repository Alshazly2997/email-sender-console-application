# Go Email Sender Console Application

A lightweight, automated CLI tool built in Go that fetches "Pending" emails from a MySQL database, dispatches them via SMTP, and updates their status to prevent duplicate sending.

## Features
* **Database Driven:** Reads email queues directly from MySQL.
* **Automated Status Tracking:** Automatically marks emails as `send` or `fail`.

## Prerequisites
- [Go](https://go.dev/doc/install) (1.18 or higher)
- [MySQL Server](https://www.mysql.com/)
- A Gmail account (or any SMTP provider)

## Setup Instructions
* **1. Database Setup:** Create a table in your MySQL database using the following schema:
  
  ```SQL
  CREATE TABLE outbox (
      id INT AUTO_INCREMENT PRIMARY KEY,
      email_address VARCHAR(255) NOT NULL,
      message_body TEXT NOT NULL,
      status ENUM('pending', 'success', 'fail') DEFAULT 'pending',
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

 **2. Environment Variables:** Create a .env file in the root directory and add your credentials
 
    ```Code snippet
    DB_PASSWORD=your_db_password
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=587
    SENDER_EMAIL=your-email@gmail.com
    SENDER_PASSWORD=your-app-password
