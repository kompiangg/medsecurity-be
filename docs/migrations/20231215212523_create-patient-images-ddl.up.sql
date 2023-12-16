CREATE TABLE IF NOT EXISTS patient_images (
  id UUID PRIMARY KEY,
  patient_id UUID NOT NULL,
  image_url TEXT NOT NULL,
  is_valid boolean NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (patient_id) REFERENCES patients (id)
);

CREATE OR REPLACE TRIGGER update_patient_images_updated_at
BEFORE UPDATE ON patient_images
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
