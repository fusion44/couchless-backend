CREATE TABLE activities
(
  id BIGSERIAL PRIMARY KEY,
  input_type VARCHAR(45), -- e.g.: 'manual', 'recorded', 'imported'
  sport_type VARCHAR(45),
  comment TEXT,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  duration int,
	boundary_north FLOAT,
	boundary_south FLOAT,
	boundary_east FLOAT,
	boundary_west FLOAT,
  time_paused INTEGER, -- seconds
  avg_pace FLOAT,
  avg_speed FLOAT,
  avg_cadence SMALLINT,
  avg_fractional_cadence SMALLINT,
  max_cadence SMALLINT,
  max_speed FLOAT,
  total_distance FLOAT,
  total_ascent INTEGER,
  total_descent INTEGER,
  max_altitude REAL,
  avg_heart_rate SMALLINT,
  max_heart_rate SMALLINT,
  total_training_effect REAL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  file_id BIGSERIAL REFERENCES file_descriptors(id) ON DELETE CASCADE,
  FOREIGN KEY (sport_type) REFERENCES sport_types(name) ON DELETE CASCADE
);