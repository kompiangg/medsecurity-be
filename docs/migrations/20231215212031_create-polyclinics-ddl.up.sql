CREATE TABLE IF NOT EXISTS polyclinics (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_patients_updated_at
BEFORE UPDATE ON polyclinics 
FOR EACH ROW 
EXECUTE PROCEDURE update_updated_at_column();