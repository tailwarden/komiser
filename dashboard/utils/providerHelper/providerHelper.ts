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
  | 'scaleway';

const providers = {
  providerLabel(arg: Provider) {
    let label;

    if (arg === 'aws') {
      label = 'Amazon Web Services';
    }

    if (arg === 'gcp') {
      label = 'Google Cloud Platform';
    }

    if (arg === 'ovh') {
      label = 'OVH';
    }

    if (arg === 'digitalocean') {
      label = 'DigitalOcean';
    }

    if (arg === 'azure') {
      label = 'Azure';
    }

    if (arg === 'tencent') {
      label = 'Tencent';
    }

    if (arg === 'civo') {
      label = 'Civo';
    }

    if (arg === 'kubernetes') {
      label = 'Kubernetes';
    }

    if (arg === 'linode') {
      label = 'Linode';
    }

    if (arg === 'oci') {
      label = 'OCI';
    }

    if (arg === 'scaleway') {
      label = 'Scaleway';
    }

    return label;
  },
  providerImg(arg: Provider) {
    let img;

    if (arg === 'aws') {
      img = '/assets/img/providers/aws.png';
    }

    if (arg === 'gcp') {
      img = '/assets/img/providers/gcp.png';
    }

    if (arg === 'ovh') {
      img = '/assets/img/providers/ovh.jpeg';
    }

    if (arg === 'digitalocean') {
      img = '/assets/img/providers/digitalocean.png';
    }

    if (arg === 'azure') {
      img = '/assets/img/providers/azure.svg';
    }

    if (arg === 'civo') {
      img = '/assets/img/providers/civo.jpeg';
    }

    if (arg === 'kubernetes') {
      img = '/assets/img/providers/kubernetes.png';
    }

    if (arg === 'linode') {
      img = '/assets/img/providers/linode.png';
    }

    if (arg === 'tencent') {
      img = '/assets/img/providers/tencent.jpeg';
    }

    if (arg === 'oci') {
      img = '/assets/img/providers/oci.png';
    }

    if (arg === 'scaleway') {
      img = '/assets/img/providers/scaleway.png';
    }

    return img;
  }
};

export default providers;
