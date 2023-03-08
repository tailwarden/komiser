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
  providerLabel(arg: Provider) {
    let label;

    if (arg.toLowerCase() === 'aws') {
      label = 'Amazon Web Services';
    }

    if (arg.toLowerCase() === 'gcp') {
      label = 'Google Cloud Platform';
    }

    if (arg.toLowerCase() === 'ovh') {
      label = 'OVH';
    }

    if (arg.toLowerCase() === 'digitalocean') {
      label = 'DigitalOcean';
    }

    if (arg.toLowerCase() === 'azure') {
      label = 'Azure';
    }

    if (arg.toLowerCase() === 'tencent') {
      label = 'Tencent';
    }

    if (arg.toLowerCase() === 'civo') {
      label = 'Civo';
    }

    if (arg.toLowerCase() === 'kubernetes') {
      label = 'Kubernetes';
    }

    if (arg.toLowerCase() === 'linode') {
      label = 'Linode';
    }

    if (arg.toLowerCase() === 'oci') {
      label = 'OCI';
    }

    if (arg.toLowerCase() === 'scaleway') {
      label = 'Scaleway';
    }

    if (arg.toLowerCase() === 'mongodbatlas') {
      label = 'MongoDB Atlas';
    }

    return label;
  },
  providerImg(arg: Provider) {
    let img;

    if (arg.toLowerCase() === 'aws') {
      img = '/assets/img/providers/aws.png';
    }

    if (arg.toLowerCase() === 'gcp') {
      img = '/assets/img/providers/gcp.png';
    }

    if (arg.toLowerCase() === 'ovh') {
      img = '/assets/img/providers/ovh.jpeg';
    }

    if (arg.toLowerCase() === 'digitalocean') {
      img = '/assets/img/providers/digitalocean.png';
    }

    if (arg.toLowerCase() === 'azure') {
      img = '/assets/img/providers/azure.svg';
    }

    if (arg.toLowerCase() === 'civo') {
      img = '/assets/img/providers/civo.jpeg';
    }

    if (arg.toLowerCase() === 'kubernetes') {
      img = '/assets/img/providers/kubernetes.png';
    }

    if (arg.toLowerCase() === 'linode') {
      img = '/assets/img/providers/linode.png';
    }

    if (arg.toLowerCase() === 'tencent') {
      img = '/assets/img/providers/tencent.jpeg';
    }

    if (arg.toLowerCase() === 'oci') {
      img = '/assets/img/providers/oci.png';
    }

    if (arg.toLowerCase() === 'scaleway') {
      img = '/assets/img/providers/scaleway.png';
    }

    if (arg.toLowerCase() === 'mongodbatlas') {
      img = '/assets/img/providers/mongodbatlas.jpg';
    }

    return img;
  }
};

export default providers;
