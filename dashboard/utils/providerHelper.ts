export type Provider =
  | 'aws'
  | 'gcp'
  | 'ovh'
  | 'digitalocean'
  | 'azure'
  | 'civo'
  | 'kubernetes'
  | 'linode'
  | 'tencent'
  | 'oci'
  | 'scaleway'
  | 'mongodbatlas';

const providers = {
  providerLabel(arg: Provider): string | undefined {
    return {
      aws: 'Amazon Web Services',
      gcp: 'Google Cloud Platform',
      ovh: 'OVH',
      digitalocean: 'DigitalOcean',
      azure: 'Azure',
      tencent: 'Tencent',
      civo: 'Civo',
      kubernetes: 'Kubernetes',
      linode: 'Linode',
      oci: 'OCI',
      scaleway: 'Scaleway',
      mongodbatlas: 'MongoDB Atlas'
    }[arg.toLowerCase()];
  },

  providerImg(arg: Provider): string | undefined {
    return {
      aws: '/assets/img/providers/aws.png',
      gcp: '/assets/img/providers/gcp.png',
      ovh: '/assets/img/providers/ovh.jpeg',
      digitalocean: '/assets/img/providers/digitalocean.png',
      azure: '/assets/img/providers/azure.svg',
      tencent: '/assets/img/providers/tencent.jpeg',
      civo: '/assets/img/providers/civo.jpeg',
      kubernetes: '/assets/img/providers/kubernetes.png',
      linode: '/assets/img/providers/linode.png',
      oci: '/assets/img/providers/oci.png',
      scaleway: '/assets/img/providers/scaleway.png',
      mongodbatlas: '/assets/img/providers/mongodbatlas.jpg'
    }[arg.toLowerCase()];
  }
};

export default providers;
