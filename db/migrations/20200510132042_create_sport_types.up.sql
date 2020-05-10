CREATE TABLE sport_types
(
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(45) NOT NULL UNIQUE,
  comment     VARCHAR(255),
  ant_id      SMALLINT UNIQUE,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--- Sport types from ANT SDK 
INSERT INTO sport_types (name, comment, ant_id) VALUES('generic', '', '0');
INSERT INTO sport_types (name, comment, ant_id) VALUES('running', '', '1');
INSERT INTO sport_types (name, comment, ant_id) VALUES('cycling', '', '2');
INSERT INTO sport_types (name, comment, ant_id) VALUES('swimming', '', '3');
INSERT INTO sport_types (name, comment, ant_id) VALUES('basketball', '', '4');
INSERT INTO sport_types (name, comment, ant_id) VALUES('soccer', '', '5');
INSERT INTO sport_types (name, comment, ant_id) VALUES('tennis', '', '6');
INSERT INTO sport_types (name, comment, ant_id) VALUES('american_football', '', '7');
INSERT INTO sport_types (name, comment, ant_id) VALUES('training', '', '8');
INSERT INTO sport_types (name, comment, ant_id) VALUES('walking', '', '9');
INSERT INTO sport_types (name, comment, ant_id) VALUES('cross_country_skiing', '', '10');
INSERT INTO sport_types (name, comment, ant_id) VALUES('alpine_skiing', '', '11');
INSERT INTO sport_types (name, comment, ant_id) VALUES('snowboarding', '', '12');
INSERT INTO sport_types (name, comment, ant_id) VALUES('rowing', '', '13');
INSERT INTO sport_types (name, comment, ant_id) VALUES('mountaineering', '', '14');
INSERT INTO sport_types (name, comment, ant_id) VALUES('hiking', '', '15');
INSERT INTO sport_types (name, comment, ant_id) VALUES('multisport', '', '16');
INSERT INTO sport_types (name, comment, ant_id) VALUES('paddling', '', '17');
INSERT INTO sport_types (name, comment, ant_id) VALUES('flying', '', '18');
INSERT INTO sport_types (name, comment, ant_id) VALUES('e_biking', '', '19');
INSERT INTO sport_types (name, comment, ant_id) VALUES('motorcycling', '', '20');
INSERT INTO sport_types (name, comment, ant_id) VALUES('boating', '', '21');
INSERT INTO sport_types (name, comment, ant_id) VALUES('driving', '', '22');
INSERT INTO sport_types (name, comment, ant_id) VALUES('golf', '', '23');
INSERT INTO sport_types (name, comment, ant_id) VALUES('hang_gliding', '', '24');
INSERT INTO sport_types (name, comment, ant_id) VALUES('horseback_riding', '', '25');
INSERT INTO sport_types (name, comment, ant_id) VALUES('hunting', '', '26');
INSERT INTO sport_types (name, comment, ant_id) VALUES('fishing', '', '27');
INSERT INTO sport_types (name, comment, ant_id) VALUES('inline_skating', '', '28');
INSERT INTO sport_types (name, comment, ant_id) VALUES('rock_climbing', '', '29');
INSERT INTO sport_types (name, comment, ant_id) VALUES('sailing', '', '30');
INSERT INTO sport_types (name, comment, ant_id) VALUES('ice_skating', '', '31');
INSERT INTO sport_types (name, comment, ant_id) VALUES('sky_diving', '', '32');
INSERT INTO sport_types (name, comment, ant_id) VALUES('snowshoeing', '', '33');
INSERT INTO sport_types (name, comment, ant_id) VALUES('snowmobiling', '', '34');
INSERT INTO sport_types (name, comment, ant_id) VALUES('stand_up_paddleboarding', '', '35');
INSERT INTO sport_types (name, comment, ant_id) VALUES('surfing', '', '36');
INSERT INTO sport_types (name, comment, ant_id) VALUES('wakeboarding', '', '37');
INSERT INTO sport_types (name, comment, ant_id) VALUES('water_skiing', '', '38');
INSERT INTO sport_types (name, comment, ant_id) VALUES('kayaking', '', '39');
INSERT INTO sport_types (name, comment, ant_id) VALUES('rafting', '', '40');
INSERT INTO sport_types (name, comment, ant_id) VALUES('windsurfing', '', '41');
INSERT INTO sport_types (name, comment, ant_id) VALUES('kitesurfing', '', '42');
INSERT INTO sport_types (name, comment, ant_id) VALUES('tactical', '', '43');
INSERT INTO sport_types (name, comment, ant_id) VALUES('jumpmaster', '', '44');
INSERT INTO sport_types (name, comment, ant_id) VALUES('boxing', '', '45');
INSERT INTO sport_types (name, comment, ant_id) VALUES('floor_climbing', '', '46');


-- Sub sport types according to the ANT SDK
INSERT INTO sport_types (name, comment, ant_id) VALUES('treadmill', 'Run/Fitness Equipment', '47');
INSERT INTO sport_types (name, comment, ant_id) VALUES('street', 'Run', '48');
INSERT INTO sport_types (name, comment, ant_id) VALUES('trail', 'Run', '49');
INSERT INTO sport_types (name, comment, ant_id) VALUES('track', 'Run', '50');
INSERT INTO sport_types (name, comment, ant_id) VALUES('spin', 'Cycling', '51');
INSERT INTO sport_types (name, comment, ant_id) VALUES('indoor_cycling', 'Cycling/Fitness Equipment', '52');
INSERT INTO sport_types (name, comment, ant_id) VALUES('road', 'Cycling', '53');
INSERT INTO sport_types (name, comment, ant_id) VALUES('mountain', 'Cycling', '54');
INSERT INTO sport_types (name, comment, ant_id) VALUES('downhill', 'Cycling', '55');
INSERT INTO sport_types (name, comment, ant_id) VALUES('recumbent', 'Cycling', '56');
INSERT INTO sport_types (name, comment, ant_id) VALUES('cyclocross', 'Cycling', '57');
INSERT INTO sport_types (name, comment, ant_id) VALUES('hand_cycling', 'Cycling', '58');
INSERT INTO sport_types (name, comment, ant_id) VALUES('track_cycling', 'Cycling', '59');
INSERT INTO sport_types (name, comment, ant_id) VALUES('indoor_rowing', 'Fitness Equipment', '60');
INSERT INTO sport_types (name, comment, ant_id) VALUES('elliptical', 'Fitness Equipment', '61');
INSERT INTO sport_types (name, comment, ant_id) VALUES('stair_climbing', 'Fitness Equipment', '62');
INSERT INTO sport_types (name, comment, ant_id) VALUES('lap_swimming', 'Swimming', '63');
INSERT INTO sport_types (name, comment, ant_id) VALUES('open_water', 'Swimming', '64');
INSERT INTO sport_types (name, comment, ant_id) VALUES('flexibility_training', 'Training', '65');
INSERT INTO sport_types (name, comment, ant_id) VALUES('strength_training', 'Training', '66');
INSERT INTO sport_types (name, comment, ant_id) VALUES('warm_up', 'Tennis', '67');
INSERT INTO sport_types (name, comment, ant_id) VALUES('match', 'Tennis', '68');
INSERT INTO sport_types (name, comment, ant_id) VALUES('exercise', 'Tennis', '69');
INSERT INTO sport_types (name, comment, ant_id) VALUES('challenge', '', '70');
INSERT INTO sport_types (name, comment, ant_id) VALUES('indoor_skiing', 'Fitness Equipment', '71');
INSERT INTO sport_types (name, comment, ant_id) VALUES('cardio_training', 'Training', '72');
INSERT INTO sport_types (name, comment, ant_id) VALUES('indoor_walking', 'Walking/Fitness Equipment', '73');
INSERT INTO sport_types (name, comment, ant_id) VALUES('e_bike_fitness', 'E-Biking', '74');
INSERT INTO sport_types (name, comment, ant_id) VALUES('bmx', 'Cycling', '75');
INSERT INTO sport_types (name, comment, ant_id) VALUES('casual_walking', 'Walking', '76');
INSERT INTO sport_types (name, comment, ant_id) VALUES('speed_walking', 'Walking', '77');
INSERT INTO sport_types (name, comment, ant_id) VALUES('bike_to_run_transition', 'Transition', '78');
INSERT INTO sport_types (name, comment, ant_id) VALUES('run_to_bike_transition', 'Transition', '79');
INSERT INTO sport_types (name, comment, ant_id) VALUES('swim_to_bike_transition', 'Transition', '80');
INSERT INTO sport_types (name, comment, ant_id) VALUES('atv', 'Motorcycling', '81');
INSERT INTO sport_types (name, comment, ant_id) VALUES('motocross', 'Motorcycling', '82');
INSERT INTO sport_types (name, comment, ant_id) VALUES('backcountry', 'Alpine Skiing/Snowboarding', '83');
INSERT INTO sport_types (name, comment, ant_id) VALUES('resort', 'Alpine Skiing/Snowboarding', '84');
INSERT INTO sport_types (name, comment, ant_id) VALUES('rc_drone', 'Flying', '85');
INSERT INTO sport_types (name, comment, ant_id) VALUES('wingsuit', 'Flying', '86');
INSERT INTO sport_types (name, comment, ant_id) VALUES('whitewater', 'Kayaking/Rafting', '87');
INSERT INTO sport_types (name, comment, ant_id) VALUES('skate_skiing', 'Cross Country Skiing', '88');
INSERT INTO sport_types (name, comment, ant_id) VALUES('yoga', 'Training', '89');
INSERT INTO sport_types (name, comment, ant_id) VALUES('pilates', 'Fitness Equipment', '90');
INSERT INTO sport_types (name, comment, ant_id) VALUES('indoor_running', 'Run', '91');
INSERT INTO sport_types (name, comment, ant_id) VALUES('gravel_cycling', 'Cycling', '92');
INSERT INTO sport_types (name, comment, ant_id) VALUES('e_bike_mountain', 'Cycling', '93');
INSERT INTO sport_types (name, comment, ant_id) VALUES('commuting', 'Cycling', '94');
INSERT INTO sport_types (name, comment, ant_id) VALUES('mixed_surface', 'Cycling', '95');
INSERT INTO sport_types (name, comment, ant_id) VALUES('navigate', '', '96');
INSERT INTO sport_types (name, comment, ant_id) VALUES('track_me', '', '97');
INSERT INTO sport_types (name, comment, ant_id) VALUES('map', '', '98');
INSERT INTO sport_types (name, comment, ant_id) VALUES('single_gas_diving', 'Diving', '99');
INSERT INTO sport_types (name, comment, ant_id) VALUES('multi_gas_diving', 'Diving', '100');
INSERT INTO sport_types (name, comment, ant_id) VALUES('gauge_diving', 'Diving', '101');
INSERT INTO sport_types (name, comment, ant_id) VALUES('apnea_diving', 'Diving', '102');
INSERT INTO sport_types (name, comment, ant_id) VALUES('apnea_hunting', 'Diving', '103');
INSERT INTO sport_types (name, comment, ant_id) VALUES('virtual_activity', '', '104');
INSERT INTO sport_types (name, comment, ant_id) VALUES('obstacle', 'Used for events where participants run, crawl through mud, climb over walls, etc.', '105');

-- Custom sport types not found in ANT SDK
INSERT INTO sport_types (name, comment, ant_id) VALUES('material_arts', '', NULL);
INSERT INTO sport_types (name, comment, ant_id) VALUES('muay_thai', 'Thai boxing', NULL);
INSERT INTO sport_types (name, comment, ant_id) VALUES('judo', '', NULL);
