import { FormEvent } from 'react';
import router from 'next/router';
import settingsService from '@services/settingsService';
import { allProviders, Provider } from './providerHelper';

export type Credentials = {};

export type AWSCredentials = Credentials & {
  source?: string;
  path?: string;
  profile?: string;
  aws_access_key_id?: string;
  aws_secret_access_key?: string;
};

export type AzureCredentials = Credentials & {
  source?: string;
  tenantId?: string;
  clientId?: string;
  clientSecret?: string;
  subscriptionId?: string;
};

export type CivoCredentials = Credentials & {
  source?: string;
  token?: string;
};

export type DigitalOceanCredentials = Credentials & {
  source?: string;
  token?: string;
};

export type KubernetesCredentials = Credentials & {
  source?: string;
  file?: string;
};

export type LinodeCredentials = Credentials & {
  token?: string;
};

export type TencentCredentials = Credentials & {
  token?: string;
};

export type ScalewayCredentials = Credentials & {
  accessKey?: string;
  secretKey?: string;
  organizationId?: string;
};

export type MongoDBAtlasCredentials = Credentials & {
  publicApiKey?: string;
  privateApiKey?: string;
  organizationId?: string;
};

export type GCPCredentials = Credentials & {
  serviceAccountKeyPath?: string;
};

export type OCICredentials = Credentials & {
  path?: string;
};

export const getPayloadFromForm = (formData: FormData, provider: Provider) => {
  const data = Object.fromEntries(formData.entries());

  switch (provider.toLocaleLowerCase()) {
    case allProviders.AWS:
      return {
        name: data.name,
        provider,
        credentials: {
          source: data.source,
          path: data.path,
          profile: data.profile,
          aws_access_key_id: data.aws_access_key_id,
          aws_secret_access_key: data.aws_secret_access_key
        }
      };
    case allProviders.AZURE:
      return {
        name: data.name,
        provider,
        credentials: {
          source: data.source,
          tenantId: data.tenantId,
          clientId: data.clientId,
          clientSecret: data.clientSecret,
          subscriptionId: data.subscriptionId
        }
      };
    case allProviders.CIVO:
      return {
        name: data.name,
        provider,
        credentials: {
          source: data.source,
          token: data.token
        }
      };
    case allProviders.DIGITAL_OCEAN:
      return {
        name: data.name,
        provider,
        credentials: {
          source: data.source,
          token: data.token
        }
      };
    case allProviders.KUBERNETES:
      return {
        name: data.name,
        provider,
        credentials: {
          source: data.source,
          file: data.file
        }
      };
    case allProviders.LINODE:
      return {
        name: data.name,
        provider,
        credentials: {
          token: data.token
        }
      };
    case allProviders.TENCENT:
      return {
        name: data.name,
        provider,
        credentials: {
          token: data.token
        }
      };
    case allProviders.SCALE_WAY:
      return {
        name: data.name,
        provider,
        credentials: {
          accessKey: data.accessKey,
          secretKey: data.secretKey,
          organizationId: data.organizationId
        }
      };
    case allProviders.MONGODB_ATLAS:
      return {
        name: data.name,
        provider,
        credentials: {
          publicApiKey: data.publicApiKey,
          privateApiKey: data.privateApiKey,
          organizationId: data.organizationId
        }
      };
    case allProviders.GCP:
      return {
        name: data.name,
        provider,
        credentials: {
          serviceAccountKeyPath: data.serviceAccountKeyPath
        }
      };
    case allProviders.OCI:
      return {
        name: data.name,
        provider,
        credentials: {
          path: data.path
        }
      };
  }
};

export const configureAccount = (
  e: FormEvent<HTMLFormElement>,
  provider: Provider,
  showToast: (toast: {
    hasError: boolean;
    title: string;
    message: string;
  }) => void,
  setHasError: (value: boolean) => void
) => {
  e.preventDefault();

  if (setHasError) setHasError(false);
  const payloadJson = JSON.stringify(
    getPayloadFromForm(new FormData(e.currentTarget), provider)
  );
  settingsService.addCloudAccount(payloadJson).then(res => {
    if (res === Error || res.error) {
      if (setHasError) setHasError(true);
    } else {
      showToast({
        hasError: false,
        title: 'Cloud account added',
        message: 'The cloud account was successfully added!'
      });
      router.push('/onboarding/cloud-accounts/');
    }
  });

  return true;
};
