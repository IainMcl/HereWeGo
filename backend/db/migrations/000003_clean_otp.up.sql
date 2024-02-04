-- CREATE A PROCEDURE TO CLEAN OTPS
-- THIS PROCEDURE WILL RUN EVERY 24 Hours
-- This will remove all rows from opts table where the current time is 
-- greater than expires_at column

-- CREATE EXTENSION pg_cron;

CREATE OR REPLACE PROCEDURE clean_otps()
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM otp WHERE expires_at < NOW();
END;
$$;

-- SCHEDULE THE CLEANING PROCEDURE
-- THIS WILL RUN EVERY 24 HOURS

-- CREATE OR REPLACE FUNCTION schedule_clean_otp()
-- RETURNS void AS $$
-- DECLARE
--   interval_seconds integer := 60 * 60 * 24;
-- BEGIN
--   PERFORM cron.schedule('0 0 * * *', 'SELECT clean_otp()', interval_seconds);
-- END;
-- $$ LANGUAGE plpgsql;

-- SELECT schedule_clean_otp();