package cassandra

const (
	createKeyspaceQuery             = "CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};"
	createClusterEndpointTableQuery = "CREATE TABLE IF NOT EXISTS %s.cluster_endpoint (cluster_id uuid, org_id uuid, cluster_name text, endpoint text, PRIMARY KEY (cluster_id, org_id));"
	createOrgClusterTableQuery      = "CREATE TABLE IF NOT EXISTS %s.org_cluster (org_id uuid, cluster_ids set<uuid>, PRIMARY KEY (org_id));"
	createAppConfigTableQuery       = "CREATE TABLE IF NOT EXISTS %s.app_config(id TEXT, created_time timestamp, last_updated_time timestamp, last_updated_user TEXT, name TEXT, chart_name TEXT, repo_name TEXT, release_name TEXT, repo_url TEXT, namespace TEXT, version TEXT, create_namespace BOOLEAN, privileged_namespace BOOLEAN, launch_ui_url TEXT, launch_ui_redirect_url TEXT, category TEXT, icon TEXT, description TEXT, launch_ui_values TEXT, override_values TEXT, PRIMARY KEY (name, version));"
)