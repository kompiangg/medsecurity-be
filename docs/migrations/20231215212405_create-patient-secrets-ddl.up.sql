CREATE TABLE IF NOT EXISTS patient_secrets (
  id UUID PRIMARY KEY,
  patient_id UUID NOT NULL,
  private_key TEXT NOT NULL,
  public_key TEXT NOT NULL,
  key_size INTEGER NOT NULL,
  salt TEXT NOT NULL,
  is_valid BOOLEAN NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (patient_id) REFERENCES patients (id)
);

CREATE OR REPLACE TRIGGER update_patient_secrets_updated_at
BEFORE UPDATE ON patient_secrets
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();