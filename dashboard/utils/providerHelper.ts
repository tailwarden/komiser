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
  TERRAFORM: 'terraform',
  PULUMI: 'pulumi'
};

export type Integration = 'slack' | 'webhook';
type IntegrationsKey = 'SLACK' | 'WEBHOOK';

export const allIntegrations: { [key in IntegrationsKey]: Integration } = {
  SLACK: 'slack',
  WEBHOOK: 'webhook'
};

export type DBProvider = 'postgres' | 'sqlite';
type DBProviderKey = 'POSTGRES' | 'SQLITE';

export const allDBProviders: { [key in DBProviderKey]: DBProvider } = {
  POSTGRES: 'postgres',
  SQLITE: 'sqlite'
};

type PlatformItem = {
  label: string;
  imgSrc: string;
};

export type Platform = {
  provider: Record<string, PlatformItem>;
  integration: Record<string, PlatformItem>;
  getImgSrc: (platformName: Provider | Integration) => string;
};

const platform: Platform = {
  provider: {
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
    terraform: {
      label: 'Terraform',
      imgSrc: '/assets/img/providers/terraform.png'
    },
    pulumi: {
      label: 'Pulumi',
      imgSrc: '/assets/img/providers/pulumi.png'
    }
  },
  integration: {
    slack: {
      label: 'Slack',
      imgSrc: '/assets/img/integrations/slack.png'
    },
    webhook: {
      label: 'Custom Web-Hook',
      imgSrc: '/assets/img/integrations/webhook.png'
    }
  },
  getImgSrc(platformName) {
    // check if img exists in '/assets/img/providers'
    if (this.provider[platformName]) return this.provider[platformName].imgSrc;

    // check if img exists in '/assets/img/integrations'
    if (this.integration[platformName])
      return this.integration[platformName].imgSrc;

    return '';
  }
};

/* search these 3:
providers
providerLabel
providerImg */
export default platform;
/* // todo:
// - refactor platform in this file
//- add Avatar component
// - write story
 - replace and refactor all instances of Avatar 
 */
