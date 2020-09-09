local config = import 'config.jsonnet';

{
  test:: config.new({
    hcloudToken: 'my-token',
    hcloudNetwork: 'my-network',
  }),

  keys: std.objectFields($.test)
}

/*
OLD EXPECTED TEST RESULT:

,
   "keys": [
      "cloud-controller-manager",
      "configMap",
      "controlPlaneServices",
      "csi.hetzner.cloud",
      "hcloud-cloud-controller-manager",
      "hcloud-csi",
      "hcloud-csi-controller",
      "hcloud-csi-controller-metrics",
      "hcloud-csi-node",
      "hcloud-csi-node-metrics",
      "hcloud-volumes",
      "hcloudSecret",
      "metrics-server",
      "metrics-server-auth-reader",
      "metrics-server:system:auth-delegator",
      "system:aggregated-metrics-reader",
      "system:cloud-controller-manager",
      "system:metrics-server",
      "v1beta1.metrics.k8s.io",
      "workarounds"
   ]
  */