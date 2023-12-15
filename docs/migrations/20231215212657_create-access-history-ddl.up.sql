CREATE TABLE IF NOT EXISTS access_histories (
    id UUID PRIMARY KEY,
    patient_id UUID,
    doctor_id UUID,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients (id) ON DELETE CASCADE,
    FOREIGN KEY (doctor_id) REFERENCES doctors (id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER update_access_history_updated_at
BEFORE UPDATE ON access_histories
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
