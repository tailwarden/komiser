export type Provider =
  | 'aws'
  | 'gcp'
  | 'digitalocean'
  | 'azure'
  | 'civo'
  | 'kubernetes'
  | 'linode'
  | 'tencent'
  | 'oci'
  | 'scaleway'
  | 'mongodbatlas'
  | 'ovh'
  | 'pulumi'
  | 'terraform';

type ProviderKey =
  | 'AWS'
  | 'GCP'
  | 'DIGITAL_OCEAN'
  | 'AZURE'
  | 'CIVO'
  | 'KUBERNETES'
  | 'LINODE'
  | 'TENCENT'
  | 'OCI'
  | 'SCALE_WAY'
  | 'MONGODB_ATLAS'
  | 'OVH'
  | 'PULUMI'
  | 'TERRAFORM';

export const allProviders: { [key in ProviderKey]: Provider } = {
  AWS: 'aws',
  GCP: 'gcp',
  DIGITAL_OCEAN: 'digitalocean',
  AZURE: 'azure',
  CIVO: 'civo',
  KUBERNETES: 'kubernetes',
  LINODE: 'linode',
  TENCENT: 'tencent',
  OCI: 'oci',
  SCALE_WAY: 'scaleway',
  MONGODB_ATLAS: 'mongodbatlas',
  OVH: 'ovh',
  TERRAFORM: 'terraform',
  PULUMI: 'pulumi'
};

export type DBProvider = 'postgres' | 'sqlite';

export enum allDBProviders {
  POSTGRES = 'postgres',
  SQLITE = 'sqlite'
}

export enum IntegrationProvider {
  SLACK = 'slack',
  WEBHOOK = 'webhook'
}

type ProviderInfo = {
  label: string;
  imgSrc: string;
};

export type Platform = {
  cloudProviders: Record<string, ProviderInfo>;
  integrationProviders: Record<string, ProviderInfo>;
  getImgSrc: (providerName: Provider | IntegrationProvider) => string;
  getLabel: (providerName: Provider | IntegrationProvider) => string;
};

const platform: Platform = {
  cloudProviders: {
    aws: {
      label: 'Amazon Web Services',
      imgSrc: '/assets/img/providers/aws.png'
    },
    gcp: {
      label: 'Google Cloud Platform',
      imgSrc: '/assets/img/providers/gcp.png'
    },
    digitalocean: {
      label: 'DigitalOcean',
      imgSrc: '/assets/img/providers/digitalocean.png'
    },
    azure: {
      label: 'Azure',
      imgSrc: '/assets/img/providers/azure.png'
    },
    civo: {
      label: 'Civo',
      imgSrc: '/assets/img/providers/civo.png'
    },
    kubernetes: {
      label: 'Kubernetes',
      imgSrc: '/assets/img/providers/kubernetes.png'
    },
    linode: {
      label: 'Linode',
      imgSrc: '/assets/img/providers/linode.png'
    },
    tencent: {
      label: 'Tencent',
      imgSrc: '/assets/img/providers/tencent.png'
    },
    oci: {
      label: 'OCI',
      imgSrc: '/assets/img/providers/oci.png'
    },
    scaleway: {
      label: 'Scaleway',
      imgSrc: '/assets/img/providers/scaleway.png'
    },
    mongodbatlas: {
      label: 'MongoDB Atlas',
      imgSrc: '/assets/img/providers/mongodbatlas.png'
    },
    ovh: {
      label: 'OVHcloud',
      imgSrc: '/assets/img/providers/ovh.png'
    },
    terraform: {
      label: 'Terraform',
      imgSrc: '/assets/img/providers/terraform.png'
    },
    pulumi: {
      label: 'Pulumi',
      imgSrc: '/assets/img/providers/pulumi.png'
    }
  },
  integrationProviders: {
    slack: {
      label: 'Slack',
      imgSrc: '/assets/img/integrations/slack.png'
    },
    webhook: {
      label: 'Custom Web-Hook',
      imgSrc: '/assets/img/integrations/webhook.png'
    }
  },

  getImgSrc(providerName) {
    const key = providerName.toLowerCase();
    if (key in this.cloudProviders) return this.cloudProviders[key].imgSrc;
    if (key in this.integrationProviders)
      return this.integrationProviders[key].imgSrc;
    return '';
  },
  getLabel(providerName) {
    const key = providerName.toLowerCase();
    if (key in this.cloudProviders) return this.cloudProviders[key].label;
    if (key in this.integrationProviders)
      return this.integrationProviders[key].label;
    return '';
  }
};

export default platform;
