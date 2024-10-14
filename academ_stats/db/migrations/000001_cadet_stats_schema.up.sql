CREATE SCHEMA IF NOT EXISTS academ_stats AUTHORIZATION postgres;

-- Permissions

GRANT ALL ON SCHEMA academ_stats TO postgres;
GRANT ALL ON SCHEMA academ_stats TO session_manager;
ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA academ_stats GRANT ALL ON SEQUENCES TO session_manager;
ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA academ_stats GRANT USAGE ON TYPES TO session_manager;
ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA academ_stats GRANT EXECUTE ON FUNCTIONS TO session_manager;
ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA academ_stats GRANT ALL ON TABLES TO session_manager;