CREATE TABLE activities
(
  id BIGSERIAL PRIMARY KEY,
  input_type VARCHAR(45), -- e.g.: 'manual', 'recorded', 'imported'
  sport_type VARCHAR(45),
  comment TEXT,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  time_paused INTEGER, -- seconds
  avg_pace FLOAT,
  avg_speed FLOAT,
  avg_cadence FLOAT,
  avg_fractional_cadence FLOAT,
  max_cadence FLOAT,
  max_speed FLOAT,
  altitude_diff_up INTEGER,
  altitude_diff_down INTEGER,
  max_altitude INTEGER,
  avg_heart_rate SMALLINT,
  max_heart_rate SMALLINT,
  total_training_effect REAL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  file_id BIGSERIAL REFERENCES file_descriptors(id) ON DELETE CASCADE,
  FOREIGN KEY (sport_type) REFERENCES sport_types(name) ON DELETE CASCADE
);