ALTER TABLE otp
DROP CONSTRAINT fk_otp_user_id,
ADD CONSTRAINT fk_otp_user_id
  FOREIGN KEY (user_id)
  REFERENCES users(id);