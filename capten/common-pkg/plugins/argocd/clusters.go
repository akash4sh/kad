package argocd

import (
	"context"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/v2/util/io"
)

func (a *ArgoCDClient) CreateCluster(ctx context.Context, clusterReq *Cluster) (*v1alpha1.Cluster, error) {
	conn, appClient, err := a.client.NewClusterClient()
	if err != nil {
		return nil, err
	}
	defer io.Close(conn)

	resp, err := appClient.Create(ctx, &cluster.ClusterCreateRequest{
		Cluster: &v1alpha1.Cluster{
			Server: clusterReq.Server,
			Name:   clusterReq.Name,
			Config: v1alpha1.ClusterConfig{
				BearerToken: clusterReq.Config.BearerToken,
				TLSClientConfig: v1alpha1.TLSClientConfig{
					ServerName: clusterReq.Config.ServerName,
					CAData:     clusterReq.Config.CAData,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArgoCDClient) DeleteCluster(ctx context.Context, clusterURL string) (*cluster.ClusterResponse, error) {
	conn, appClient, err := a.client.NewClusterClient()
	if err != nil {
		return nil, err
	}
	defer io.Close(conn)

	resp, err := appClient.Delete(ctx, &cluster.ClusterQuery{
		Id: &cluster.ClusterID{
			Value: clusterURL,
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ArgoCDClient) GetCluster(ctx context.Context, clusterURL string) (*v1alpha1.Cluster, error) {
	conn, appClient, err := a.client.NewClusterClient()
	if err != nil {
		return nil, err
	}
	defer io.Close(conn)

	cluster, err := appClient.Get(ctx, &cluster.ClusterQuery{
		Id: &cluster.ClusterID{
			Value: clusterURL,
		},
	})
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (a *ArgoCDClient) ListClusters(ctx context.Context) (*v1alpha1.ClusterList, error) {
	conn, appClient, err := a.client.NewClusterClient()
	if err != nil {
		return nil, err
	}
	defer io.Close(conn)

	list, err := appClient.List(ctx, &cluster.ClusterQuery{})
	if err != nil {
		return nil, err
	}

	return list, nil
}
