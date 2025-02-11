CREATE TABLE file_descriptors
(
  id BIGSERIAL PRIMARY KEY,
  file_name VARCHAR(255),
  content_type VARCHAR(45),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL
);