CREATE TABLE IF NOT EXISTS cluster_app_config (
    release_name varchar(80) PRIMARY KEY,
    app_name varchar(80),
    plugin_name varchar(80),
    plugin_store_type INTEGER,
    category varchar(40),
    description TEXT,
    repo_url varchar(200),
    version varchar(40),
    namespace varchar(80),
    ui_endpoint varchar(200),
    api_endpoint varchar(200),
    ui_module_endpoint varchar(200),
    default_app BOOLEAN,
    privileged_namespace BOOLEAN,
    install_status varchar(40),
    icon BYTEA,
    override_values TEXT,
    launch_ui_values TEXT,
    template_values TEXT,
    last_update_time TIMESTAMP
);